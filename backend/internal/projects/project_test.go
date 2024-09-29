package projects

import (
	"context"
	"testing"

	"node-nanny/backend/internal/configs"
	"node-nanny/backend/internal/service"
)

func TestSaveProject(t *testing.T) {
	db := service.InitConnection()
	defer db.Close()

	cfg := configs.Configs{}
	projectStore := NewProjectSqlStore(db, cfg.ProjectSqlTable())
	projectSvc := NewService(projectStore)

	project := Project{
		Name:         "test_project",
		Path:         "$home/Dev/",
		Detail:       "23-23-aa",
		Dependencies: "23-23-aa",
		Subproject:   "23-23-aa",
		Files:        "23-23-aa",
		Logs:         "23-23-aa",
	}

	ctx := context.Background()
	id, err := projectSvc.SaveProject(ctx, &project)
	if err != nil || id == "" {
		t.Errorf("Failed: %s", err)
	}
}
