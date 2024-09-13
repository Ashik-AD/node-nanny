package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
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
	IsDev   bool
}

type EnvVariable struct {
	Name, Value string
}

type Project struct {
	Path string
	Name string

	Files        []ProjectFile
	Dependencies []ProjectDependency
	EnvVariables map[string]string
	SubProjects  []Project
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

	dirs, dirErr := os.ReadDir(selection)
	if dirErr != nil {
		log.Fatal(dirErr)
	}

	project = Project{}

	for _, obj := range dirs {
		if obj.IsDir() {
			_, ok := DIRECTORIES_TO_BE_IGNORE[strings.ToLower(obj.Name())]

			if !ok {
				fmt.Println("Scanning Dir: ", obj.Name())
				objAbspath := path.Join(selection, obj.Name())

				if childDir, dirErr := os.ReadDir(objAbspath); dirErr != nil {
					log.Fatal(dirErr)
					log.Fatal("Failed scan child directory %q in %v", objAbspath, selection)
				} else {

					isProject := slices.ContainsFunc(childDir, func(entry fs.DirEntry) bool {
						return entry.Name() == "package.json"
					})
					projectFiles := []ProjectFile{}

					for _, file := range childDir {
						if !file.IsDir() {
							fileName := file.Name()
							if isProject {
								if hasAllowedFile(ALLOWED_FILES, fileName) {
									file := ProjectFile{
										Path: path.Join(objAbspath, fileName),
										Name: fileName,
									}
									projectFiles = append(projectFiles, file)
									fmt.Printf("\n\tFound file: %v in path: %v\n", fileName, objAbspath)
								}
							}
						}
					}

					if isProject {
						childProject := Project{
							Name:  obj.Name(),
							Path:  objAbspath,
							Files: projectFiles,
						}
						project.SubProjects = append(project.SubProjects, childProject)
						isProject = false
						projectFiles = []ProjectFile{}
					}
				}
			}
		} else {
			if !obj.IsDir() {

				fileName := obj.Name()
				isProject := slices.ContainsFunc(dirs, func(entry fs.DirEntry) bool {
					return entry.Name() == "package.json"
				})

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
	}

	if !reflect.ValueOf(project).IsZero() {
		// runtime.EventsEmit(a.ctx, "onDirectoryScanDone", "project directory scanning done")

		dependencies := []ProjectDependency{}
		envVariables := map[string]string{}

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
				depList, ok := data["dependencies"]
				if ok {
					for key, value := range depList.(map[string]interface{}) {
						dep := ProjectDependency{key, value.(string), false}
						dependencies = append(dependencies, dep)
					}
				} else {
					fmt.Println("key is not found")
				}

				devDepList, ok := data["devDependencies"]
				if ok {
					for key, value := range devDepList.(map[string]interface{}) {
						dep := ProjectDependency{key, value.(string), true}
						dependencies = append(dependencies, dep)
					}
				}

			case ".env":
				//Todo implement
				scanner := bufio.NewScanner(file)

				for scanner.Scan() {
					if scanner.Text() != "" {
						values := strings.SplitN(scanner.Text(), "=", 2)
						envVariables[values[0]] = values[1]
					}
				}
			}
		}

		// Checking dependencies on sub project
		for i, sub := range project.SubProjects {
			subDependencies := []ProjectDependency{}
			subEnvVariables := map[string]string{}

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

					depList, ok := data["dependencies"]
					if ok {
						for key, value := range depList.(map[string]interface{}) {
							dep := ProjectDependency{key, value.(string), false}
							subDependencies = append(subDependencies, dep)
						}
					} else {
						fmt.Println("key is not found")
					}

					devDepList, ok := data["devDependencies"]
					if ok {
						for key, value := range devDepList.(map[string]interface{}) {
							dep := ProjectDependency{key, value.(string), true}
							subDependencies = append(subDependencies, dep)
						}
					}
					project.SubProjects[i].Dependencies = subDependencies
				case ".env":
					scanner := bufio.NewScanner(openedFile)

					for scanner.Scan() {
						if scanner.Text() != "" {
							values := strings.SplitN(scanner.Text(), "=", 2)
							subEnvVariables[values[0]] = values[1]
						}
					}
					project.SubProjects[i].EnvVariables = subEnvVariables
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
