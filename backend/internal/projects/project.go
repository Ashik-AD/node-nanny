package projects

import (
	"context"
	"errors"
	"strings"
)

type store interface {
	SaveProject(ctx context.Context, project *Project) (string, error)
}

type Project struct {
	ID           string
	Name         string
	Path         string
	Detail       string
	Dependencies string
	Files         string
	Subproject   string
	Logs         string
}

type Projects struct {
	store store
}

func (project *Project) Senitize() {
	project.Name = strings.TrimSpace(project.Name)
	project.Path = strings.TrimSpace(project.Path)
}

func (project *Project) ValidateForCreate() error {
	if project == nil {
		return errors.New("empty project")
	}

	project.Senitize()
	if project.Name == "" {
		return errors.New("project name can't be empty")
	}
	if project.Path == "" {
		return errors.New("project path can't be empty")
	}
	return nil
}

func (pr *Projects) SaveProject(ctx context.Context, project *Project) (string, error) {
	project.Senitize()
	err := project.ValidateForCreate()
	if err != nil {
		return "", err
	}
	pid, err := pr.store.SaveProject(ctx, project)
	if err != nil {
		return "", err
	}

	return pid, err
}

func NewService(store store) *Projects {
	return &Projects{
		store: store,
	}
}
