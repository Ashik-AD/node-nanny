package files

import (
	"fmt"
	"testing"

	"node-nanny/backend/internal/configs"
	"node-nanny/backend/internal/service"
)

func TestFiles(t *testing.T) {
	db := service.InitConnection()
	cfg := configs.Configs{}
	store := NewFileSqlStore(db, cfg.FilesSqlTable())
	svc := NewFilesService(store)

	t.Run("get file list by pid", func(t *testing.T) {
		list, err := svc.GetFilesByProject("d0f3e4c2-5e9f-4f39-9d1d-3d983e2c4b0a")
		if err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Println("files", list)
		}
	})

	t.Run("get file by ID", func(t *testing.T) {
		file, err := svc.GetFileByID("298e3c8c-e734-47e1-9d78-5be69f7d0d8b")
		if err != nil {
			t.Errorf("%v", err)
		} else {
			fmt.Println(file)
		}
	})
}
