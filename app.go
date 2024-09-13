package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"slices"
	"strings"

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

type Project struct {
	Path string
	Name string

	Files           []ProjectFile
	Dependencies    []ProjectDependency
	DevDependencies []ProjectDependency
	EnvVariables    map[string]string
	SubProjects     []Project
}

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

func (a *App) OpenProjectSelect() (project Project) {
	DIRECTORIES_TO_BE_IGNORE := map[string]bool{".git": false, "node_modules": false, "dist": false, "config": false, "util": false, "build": false}
	ALLOWED_FILES := []string{"package.json", ".env", "readme.md", "package-lock.json"}

	// Open native file explorer
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{Title: "Select Project Directory"})

	if err != nil {
		log.Fatal("Failed to open file explorer ", err)
	}

	// Get child/sub directories and files
	dirs, dirErr := os.ReadDir(selection)
	if dirErr != nil {
		log.Fatal(dirErr)
	}

	project = Project{}
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
					}
				}
			}
		} else {
			// check parent/root directory is a project
			isProject := slices.ContainsFunc(dirs, func(entry os.DirEntry) bool {
				return entry.Name() == "package.json"
			})

			fileName := fs.Name()
			if isProject {
				if hasAllowedFile(ALLOWED_FILES, fileName) {
					file := ProjectFile{
						Path: path.Join(selection, fileName),
						Name: fileName,
					}
					project.Files = append(project.Files, file)
				}
			}
		}
	}

	// if project is not empty, then
	// analyze project dependencies, env. variables etc
	if !reflect.ValueOf(project).IsZero() {
		// runtime.EventsEmit(a.ctx, "onDirectoryScanDone", "project directory scanning done")

		dependencies := []ProjectDependency{} // contains both dev. and prod. dependencies
		envVariables := map[string]string{}   // containes env. variables

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

		// Checking dependencies on sub project
		if len(project.SubProjects) > 0 {
			for i, sub := range project.SubProjects {
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

		project.Name = path.Base(selection)
		project.Path = selection
		project.Dependencies = dependencies
		project.EnvVariables = envVariables
	}

	fmt.Println(path.Base(selection))
	return project
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
		fmt.Println("key `dependecies` is not found")
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
