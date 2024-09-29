package projects

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type sqlstore struct {
	dbd       *sql.DB
	tablename string
}

func (ss *sqlstore) SaveProject(ctx context.Context, project *Project) (string, error) {
	project.ID = ss.NewProjectID()

	query := fmt.Sprintf(`INSERT INTO %s (id, name, path, detail, files, 
                dependencies, subprojects, logs) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		ss.tablename)
	stm, err := ss.dbd.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	_, serr := stm.ExecContext(ctx,
		&project.ID, &project.Name,
		&project.Path, &project.Detail,
		&project.Files, &project.Dependencies,
		&project.Subproject, &project.Logs,
	)
	if serr != nil {
		return "", fmt.Errorf("failed to insert project: %s", err)
	}
	return project.ID, nil
}

func (ss *sqlstore) NewProjectID() string {
	return uuid.NewString()
}

func NewProjectSqlStore(db *sql.DB, tablename string) store {
	return &sqlstore{db, tablename}
}
