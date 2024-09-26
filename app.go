package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type SelectorOption struct {
	Title string
}

type ProjectFile struct {
	Path string
	Name string
}

type ProjectDependency struct {
	Name    string
	Version string
}

type EnvVariable struct {
	Name, Value string
}

type Progress struct {
	Name       string
	Status     string
	Additional []string
}

type PackageDetails struct {
	Name       string
	LatestVer  string
	CurrentVer string
	Homepage   string
	Npm        string
	Repository string
}

type Project struct {
	Path string
	Name string

	Files           []ProjectFile
	Dependencies    []ProjectDependency
	DevDependencies []ProjectDependency
	EnvVariables    map[string]string
	SubProjects     []Project
	Miscellaneous   struct {
		TotalModule int
		Disk        int64
	}
}

type ProjectDependencyDetails struct {
	Name       string
	CurrentVer string
	LatestVer  string
	Npm        string
	Hompeage   string
	Repository string
}

type ProjectDetails struct {
	Path string
	Name string

	Files           []ProjectFile
	SubProjects     []ProjectDetails
	Dependencies    []PackageDetails
	DevDependencies []PackageDetails
	Miscellaneous   struct {
		TotalModule int
		Disk        int64
	}
}

const (
	NPM_RESGISTRY_URL = "https://registry.npmjs.com/-/v1"
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given fileName
func (a *App) Greet(fileName string) string {
	return fmt.Sprintf("Hello %s, It's show time!", fileName)
}

func (a *App) OpenProjectSelect() (projectDetails ProjectDetails) {
	DIRECTORIES_TO_BE_IGNORE := map[string]bool{".git": false, "node_modules": false, "dist": false, "config": false, "util": false, "build": false, "lib": false, "utils": false}
	ALLOWED_FILES := []string{"package.json", ".env", "readme.md", "package-lock.json"}

	progress := []Progress{
		Progress{Name: "Checking project", Status: "pending"},
		Progress{Name: "Checking project dependecies", Status: "pending"},
		Progress{Name: "Preparing project", Status: "pending"},
	}

	project := Project{}

	// Open native file explorer
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{Title: "Select Project Directory"})

	if err != nil {
		log.Fatal("Failed to open file explorer ", err)
	}

	// read selected directory
	dirs, dirErr := os.ReadDir(selection)
	if dirErr != nil {
		log.Println(dirErr)
		return
	}

	// start project scan progress
	progress[0].Status = "scanning"
	runtime.EventsEmit(a.ctx, "start", progress)

	for _, fs := range dirs {
		// check if the fs is directory, then it can be possibly sub-project
		if fs.IsDir() {

			// ignore some of the directories
			_, ok := DIRECTORIES_TO_BE_IGNORE[strings.ToLower(fs.Name())]
			if !ok {
				fmt.Println("Scanning Dir: ", fs.Name())

				// Absolute path the directory/file
				fsAbsPath := path.Join(selection, fs.Name())

				if childDir, dirErr := os.ReadDir(fsAbsPath); dirErr != nil {
					log.Fatal(dirErr)
					log.Fatal("Failed scan child directory %q in %v", fsAbsPath, selection)
				} else {
					// If the directory contain `package.json` file
					// Then treat directory as sub|child-project
					isProject := slices.ContainsFunc(childDir, func(entry os.DirEntry) bool {
						return entry.Name() == "package.json"
					})

					projectFiles := []ProjectFile{}

					// ignore dirctories in the sub-project
					// if sub-directry is project directory,
					// then add to project files
					for _, file := range childDir {
						if !file.IsDir() {
							fileName := file.Name()
							if isProject {
								if hasAllowedFile(ALLOWED_FILES, fileName) {
									file := ProjectFile{
										Path: path.Join(fsAbsPath, fileName),
										Name: fileName,
									}
									projectFiles = append(projectFiles, file)
									fmt.Printf("\n\tFound file: %v in path: %v\n", fileName, fsAbsPath)
								}
							}
						}
					}

					if isProject {
						childProject := Project{
							Name:  fs.Name(),
							Path:  fsAbsPath,
							Files: projectFiles,
						}
						project.SubProjects = append(project.SubProjects, childProject)
						isProject = false
						projectFiles = []ProjectFile{}

						// add subproject name into Progress
						progress[0].Additional = append(progress[0].Additional, fmt.Sprintf("Found sub project `%s`", childProject.Name))
						runtime.EventsEmit(a.ctx, "scanUpdate", progress)
					}
				}
			}
		} else {
			// check parent/root directory is a project
			isProject := slices.ContainsFunc(dirs, func(entry os.DirEntry) bool {
				return entry.Name() == "package.json"
			})

			if isProject {
				fileName := fs.Name()
				if hasAllowedFile(ALLOWED_FILES, fileName) {
					file := ProjectFile{
						Path: path.Join(selection, fileName),
						Name: fileName,
					}
					project.Files = append(project.Files, file)
				}

				progress[0].Status = "completed"
				runtime.EventsEmit(a.ctx, "scanUpdate", progress)
			}
		}
	}

	log.Println(project.SubProjects)
	// if project is not empty, then
	// analyze project dependencies, env. variables etc
	if !reflect.ValueOf(project).IsZero() {
		progress[1].Status = "scanning"
		runtime.EventsEmit(a.ctx, "checkingDependency", progress)

		dependencies := []ProjectDependency{} // contains both dev. and prod. dependencies
		envVariables := map[string]string{}   // containes env. variables

		if len(project.Files) > 0 {
			progress[1].Additional = append(progress[1].Additional, "Checking dependencies on root project")
			runtime.EventsEmit(a.ctx, "checkingDependency", progress)
			for _, fileObj := range project.Files {

				filePath := fileObj.Path
				fileName := fileObj.Name

				file, err := os.Open(filePath)
				if err != nil {
					log.Fatal("Failed to open file `", fileName, "`")
					log.Fatal(err)
				}
				defer file.Close()

				content, err := os.ReadFile(filePath)
				if err != nil {
					log.Fatal("Failed to read file `", fileName)
				}

				switch strings.ToLower(fileName) {
				case "package.json":
					var data map[string]interface{}

					err := json.Unmarshal(content, &data)
					if err != nil {
						log.Fatal("Failed to parse `", fileName, "`")
					}
					prod, dev := getProjectDependencies(data)
					project.Dependencies = prod
					project.DevDependencies = dev

				case ".env":
					envVariables = getEnvVariables(file)
				}
			}
		}

		// Checking dependencies on sub project
		if len(project.SubProjects) > 0 {
			log.Println("Checking dependencies on sub project directory")
			for i, sub := range project.SubProjects {

				progress[1].Additional = append(progress[1].Additional, fmt.Sprintf("Checking dependencies on `%s`", sub.Name))
				runtime.EventsEmit(a.ctx, "checkingDependency", progress)

				for _, file := range sub.Files {
					openedFile, err := os.Open(file.Path)
					if err != nil {
						log.Fatalf("\n\tFailed to open file `%v` at path: %v", file.Name, file.Path)
						log.Fatal(err)
					}
					defer openedFile.Close()

					content, err := os.ReadFile(file.Path)
					if err != nil {
						log.Fatalf("\n\tFailed to read file `%v` at path: %v", file.Name, file.Path)
						log.Fatal(err)
					}

					switch strings.ToLower(file.Name) {
					case "package.json":
						var data map[string]interface{}

						jsonErr := json.Unmarshal(content, &data)
						if jsonErr != nil {
							log.Fatal("Failed to parse JSON of file: ", file.Name)
							log.Fatal(jsonErr)
						}
						prod, dev := getProjectDependencies(data)
						project.SubProjects[i].Dependencies = prod
						project.SubProjects[i].DevDependencies = dev

					case ".env":
						subEnvVaribles := getEnvVariables(openedFile)
						project.SubProjects[i].EnvVariables = subEnvVaribles
					}
				}
			}
		}

		log.Println("Finished dependency checking")
		progress[1].Status = "completed"
		runtime.EventsEmit(a.ctx, "checkingDependency", progress)

    //Sleep for 400ms before start
		time.Sleep(time.Millisecond*400)

		progress[2].Status = "scanning"
		runtime.EventsEmit(a.ctx, "preparingProject", progress)

		// calculate size of the directory
		totalSize, err := getDirctorySize(selection)
		if err != nil {
			log.Println(err)
		}
		projectDetails.Miscellaneous.Disk = totalSize

		// get total node modules
		totalNodeModule, err := countNodeModules(path.Join(selection, "node_modules"))
		if err != nil {
			log.Fatal(err)
		}
		projectDetails.Miscellaneous.TotalModule = totalNodeModule

		// disk usage and total node modules in sub directory project
		if len(project.SubProjects) > 0 {
			for idx, subproject := range project.SubProjects {
				dirPath := path.Join(subproject.Path, "node_modules")

				// count node modules
				count, err := countNodeModules(dirPath)
				if err != nil {
					log.Printf("Warning: `%s` => %s\n", err, dirPath)
				}
				project.SubProjects[idx].Miscellaneous.TotalModule = count
				totalNodeModule = totalNodeModule + count

				// get the size of the directory
				size, err := getDirctorySize(subproject.Name)
				if err != nil {
					log.Printf("Warning: `%s` in `%s`\n", err, subproject.Name)
				}
				project.SubProjects[idx].Miscellaneous.Disk = size
				totalSize = totalSize + size
			}
			projectDetails.Miscellaneous.TotalModule = totalNodeModule
			projectDetails.Miscellaneous.Disk = totalSize
		}

		// get package details from npm registry
		parentPkg := [][]ProjectDependency{project.Dependencies, project.DevDependencies}
		if len(parentPkg) > 0 {
			for idx, pkgs := range parentPkg {
				detailArr, err := concurrentPackageIterate(&pkgs)
				if err != nil {
					log.Println("Failed concurrent packge: ", err)
				} else {
					if idx == 0 {
						projectDetails.Dependencies = detailArr
					} else {
						projectDetails.DevDependencies = detailArr
					}
				}
			}
		}

		// get packge details of the sub project from npm registry
		if len(project.SubProjects) > 0 {
			projectDetails.SubProjects = make([]ProjectDetails, len(project.SubProjects))
			for pidx, pr := range project.SubProjects {
				deps := [][]ProjectDependency{pr.Dependencies, pr.DevDependencies}
				if len(deps) > 0 {
					for idx, pkgs := range deps {
						detailArr, err := concurrentPackageIterate(&pkgs)
						if err != nil {
							log.Println("Failed concurrent package in child directory: ", err)
						} else {
							if idx == 0 {
								projectDetails.SubProjects[pidx].Dependencies = detailArr
							} else {
								projectDetails.SubProjects[pidx].DevDependencies = detailArr
							}
						}
					}
				}
			}
		}

		projectDetails.Name = path.Base(selection)
		projectDetails.Path = selection
		project.Dependencies = dependencies
		project.EnvVariables = envVariables

    time.Sleep(time.Millisecond*400)
    progress[2].Status = "completed"
    runtime.EventsEmit(a.ctx, "preparingProject", progress)
	}
	log.Println(projectDetails)
	return projectDetails
}

func hasAllowedFile(FilesName []string, Name string) bool {
	isExist := false
	for _, file := range FilesName {
		if strings.ToLower(file) == strings.ToLower(Name) {
			isExist = true
			break
		}
	}
	return isExist
}

// return dependencies, and devDependencies from json
func getProjectDependencies(jsonMap map[string]interface{}) (prod, dev []ProjectDependency) {
	prodDepList, ok := jsonMap["dependencies"]
	if !ok {
		fmt.Println("key `dependencies` is not found")
	} else {
		for key, value := range prodDepList.(map[string]interface{}) {
			dep := ProjectDependency{key, value.(string)}
			prod = append(prod, dep)
		}
	}

	devDepList, ok := jsonMap["devDependencies"]
	if !ok {
		fmt.Println("key `devDependencies` is not found")
	} else {
		for key, value := range devDepList.(map[string]interface{}) {
			dep := ProjectDependency{key, value.(string)}
			dev = append(dev, dep)
		}
	}
	return prod, dev
}

// return env. variables from .env file
func getEnvVariables(envFile *os.File) map[string]string {
	envMap := map[string]string{}
	scanner := bufio.NewScanner(envFile)

	defer envFile.Close()

	for scanner.Scan() {
		if scanner.Text() != "" {
			values := strings.SplitN(scanner.Text(), "=", 2)
			envMap[values[0]] = values[1]
		}
	}
	return envMap
}

func getDirctorySize(dirPath string) (int64, error) {
	var totalSize int64
	totalSize = 0

	err := filepath.Walk(dirPath, func(_ string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			totalSize = totalSize + info.Size()
		}
		return nil
	})

	return totalSize, err
}

func countNodeModules(node_path string) (int, error) {
	totalModule := 0

	dirs, err := os.ReadDir(node_path)
	if err != nil {
		return 0, err
	}

	for _, file := range dirs {
		if file.IsDir() && !strings.Contains(file.Name(), ".") {
			totalModule = totalModule + 1
		}
	}
	return totalModule, nil
}

// Get packge details from the endpoint: https://registry.npmjs.org/-/v1/search?text=?&size=1
func getPackageDetails(pkgName string) (PackageDetails, error) {
	var jsonData map[string]interface{}
	pkg := PackageDetails{}

	query := fmt.Sprintf("%s/search?text=%s&size=1", NPM_RESGISTRY_URL, pkgName)
	res, err := http.Get(query)
	if err != nil {
		log.Println(err)
		return pkg, err
	}

	if res.StatusCode != http.StatusOK {
		return pkg, fmt.Errorf("Something went wrong while fetching packge details")
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return pkg, err
	}

	if body == nil || len(body) <= 0 {
		return pkg, fmt.Errorf("%s not found")
	}

	jsonErr := json.Unmarshal(body, &jsonData)
	if jsonErr != nil {
		return pkg, jsonErr
	}
	log.Println(jsonData)
	d := jsonData["objects"].([]interface{})[0].(map[string]interface{})["package"].(map[string]interface{})

	pkg.Name = pkgName
	pkg.LatestVer = d["version"].(string)
	if d["links"] != nil {
		if npmLink, ok := d["links"].(map[string]interface{})["npm"].(string); ok {
			pkg.Npm = npmLink
		}
		if homeLink, ok := d["links"].(map[string]interface{})["homepage"].(string); ok {
			pkg.Homepage = homeLink
		}
		if repoLink, ok := d["links"].(map[string]interface{})["repository"].(string); ok {
			pkg.Repository = repoLink
		}
	}

	return pkg, nil
}

func concurrentPackageIterate(pkgs *[]ProjectDependency) ([]PackageDetails, error) {
	result := []PackageDetails{}
	resChannel := make(chan PackageDetails)
	var error error

	for _, pkg := range *pkgs {
		go func() {
			res, err := getPackageDetails(pkg.Name)
			if err != nil {
				log.Println("Failed to fetch package details: ", err)
				error = err
				return
			}
			res.CurrentVer = pkg.Version
			resChannel <- res
		}()

		if error != nil {
			break
		}
	}

	if error != nil {
		return nil, error
	}

	for i := 0; i < len(resChannel); i++ {
		r := <-resChannel
		result[i] = r
	}

	return result, nil
}
