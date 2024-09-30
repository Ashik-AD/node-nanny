package dependencies

import (
	"fmt"
	"testing"

	"node-nanny/backend/internal/configs"
	"node-nanny/backend/internal/service"
)

func TestDependency(t *testing.T) {
	db := service.InitConnection()
	cfg := configs.Configs{}
	dependencyStore := NewDependencyStore(db, cfg.DependencySqlTable(), cfg.PackgeDetailsSqlTable())
	dpSvc := NewDependencySvc(dependencyStore)

	t.Run("fetch dependencies", func(t *testing.T) {
		list, err := dpSvc.GetDependencies("d0f3e4c2-5e9f-4f39-9d1d-3d983e2c4b0a")
		if err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Println(list)
		}
	})

	t.Run("fetch dependency by ID", func(t *testing.T) {
		dep, err := dpSvc.GetDependency("d0f3e4c2-5e9f-4f39-9d1d-3d983e2c4b0a", "d9b5a5d4-3456-4b6f-8fd2-54c4b2d1f1db")
		if err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Println(dep)
		}
	})
}
