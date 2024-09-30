package subprojects

import (
	"fmt"
	"testing"

	"node-nanny/backend/internal/configs"
	"node-nanny/backend/internal/service"
)

func TestDependency(t *testing.T) {
	db := service.InitConnection()
	cfg := configs.Configs{}
	store := NewSubprojectSqlStore(db, cfg.SubprojectSqlTable())
	svc := NewSubprojectService(store)

	t.Run("fetch subproject list By parent project id(root_id)", func(t *testing.T) {
		list, err := svc.GetSubprojectsByParentID("d0f3e4c2-5e9f-4f39-9d1d-3d983e2c4b0a")
		if err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Println(list)
		}
	})

	t.Run("fetch subproect by ID", func(t *testing.T) {
		dep, err := svc.GetSubprojectByID("3d569b07-8b54-4746-bab4-4d82f6b0f1c8")
		if err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Println(dep)
		}
	})
}
