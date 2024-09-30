package files

import (
	"errors"
	"strings"
)

type File struct {
	ID   string
	Name string
	Path string
	PID  string
}

type store interface {
	SaveFiles(files *[]File) (*[]File, error)
	GetFilesByProject(pid string) (*[]File, error)
	GetFileByID(id string) (*File, error)
}

type Files struct {
	store store
}

func (File *File) Senitize() {
	File.PID = trimSpace(File.PID)
	File.Name = trimSpace(File.Name)
	File.Path = trimSpace(File.Path)
}

func (File *File) ValidateForCreate() error {
	if File == nil {
		return errors.New("empty File")
	}

	File.Senitize()
	if File.PID == "" {
		return errors.New("project id is can't be empty")
	}
	if File.Name == "" {
		return errors.New("File name can't be empty")
	}
	if File.Path == "" {
		return errors.New("File path can't be empty")
	}
	return nil
}

func (fs *Files) SaveFiles(files *[]File) (*[]File, error) {
	var err error
	for _, file := range *files {
		err = file.ValidateForCreate()
		if err != nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	files, err = fs.store.SaveFiles(files)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (fs *Files) GetFilesByProject(pid string) (*[]File, error) {
	if trimSpace(pid) == "" {
		return nil, errors.New("pid can't be empty")
	}
	files, err := fs.store.GetFilesByProject(pid)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (fs *Files) GetFileByID(id string) (*File, error) {
	if trimSpace(id) == "" {
		return nil, errors.New("id can't be empty")
	}
	file, err := fs.store.GetFileByID(id)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func trimSpace(str string) string {
	return strings.TrimSpace(str)
}

func NewFilesService(store store) *Files {
	return &Files{store: store}
}
