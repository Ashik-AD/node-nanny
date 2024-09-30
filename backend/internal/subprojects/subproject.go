package subprojects

import (
	"context"
	"errors"
	"strings"
)

type store interface {
	SaveSubproject(ctx context.Context, project *Subproject) (string, error)
	GetSubprojectsByParentID(root_id string) (*[]Subproject, error)
	GetSubprojectByID(pid string) (*Subproject, error)
}

type Subproject struct {
	ID           string
	RootID       string
	Name         string
	Path         string
	Detail       string
	Dependencies string
	Files        string
	Logs         string
}

type Subprojects struct {
	store store
}

func (project *Subproject) Senitize() {
	project.Name = strings.TrimSpace(project.Name)
	project.Path = strings.TrimSpace(project.Path)
}

func (project *Subproject) ValidateForCreate() error {
	if project == nil {
		return errors.New("empty subproject")
	}

	project.Senitize()
	if project.RootID == "" {
		return errors.New("subproject root_id can't be empty")
	}
	if project.Name == "" {
		return errors.New("subproject name can't be empty")
	}
	if project.Path == "" {
		return errors.New("subproject path can't be empty")
	}
	return nil
}

func (pr *Subprojects) SaveSubproject(ctx context.Context, project *Subproject) (string, error) {
	project.Senitize()
	err := project.ValidateForCreate()
	if err != nil {
		return "", err
	}
	pid, err := pr.store.SaveSubproject(ctx, project)
	if err != nil {
		return "", err
	}

	return pid, err
}

func (pr *Subprojects) GetSubprojectByID(pid string) (*Subproject, error) {
	if strings.TrimSpace(pid) == "" {
		return nil, errors.New("pid is missing")
	}

	sp, err := pr.store.GetSubprojectByID(pid)
	if err != nil {
		return nil, err
	}
	return sp, nil
}

func (pr *Subprojects) GetSubprojectsByParentID(root_id string) (*[]Subproject, error) {
	if strings.TrimSpace(root_id) == "" {
		return nil, errors.New("root_id is missing")
	}

	prList, err := pr.store.GetSubprojectsByParentID(root_id)
	if err != nil {
		return nil, err
	}
	return prList, nil
}

func NewSubprojectService(store store) *Subprojects {
	return &Subprojects{
		store: store,
	}
}
