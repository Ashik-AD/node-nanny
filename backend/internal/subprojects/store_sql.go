package subprojects

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

type sqlstore struct {
	dbd       *sql.DB
	tablename string
}

func (ss *sqlstore) SaveSubproject(ctx context.Context, project *Subproject) (string, error) {
	project.ID = ss.NewProjectID()

	query := fmt.Sprintf(`
        INSERT INTO %s (id, root_id, name, path, detail, files, 
        dependencies, logs) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		ss.tablename)

	stm, err := ss.dbd.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stm.Close()

	_, serr := stm.ExecContext(ctx,
		&project.ID, &project.RootID, &project.Name,
		&project.Path, &project.Detail,
		&project.Files, &project.Dependencies,
		&project.Logs,
	)
	if serr != nil {
		return "", fmt.Errorf("failed to insert project: %s", err)
	}
	return project.ID, nil
}

func (ss *sqlstore) GetSubprojectsByParentID(root_id string) (*[]Subproject, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE root_id = ?", ss.tablename)
	stm, err := ss.dbd.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	rows, err := stm.Query(root_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Subproject
	for rows.Next() {
		var pr Subproject
		if err := rows.Scan(
			&pr.ID, &pr.RootID, &pr.Name, &pr.Path,
			&pr.Detail, &pr.Files,
			&pr.Dependencies,
			&pr.Logs,
		); err != nil {
			return nil, err
		}
		projects = append(projects, pr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &projects, nil
}

func (ss *sqlstore) GetSubprojectByID(pid string) (*Subproject, error) {
	var pr Subproject

	pid = strings.TrimSpace(pid)
	if pid == "" {
		return nil, errors.New("empty project ID")
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", ss.tablename)
	row := ss.dbd.QueryRow(query, pid)
	if err := row.Scan(&pr.ID, &pr.RootID, &pr.Name, &pr.Path, &pr.Detail, &pr.Files, &pr.Dependencies, &pr.Logs); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No record found with the ID: %s", pid)
		}
		return nil, err
	}

	return &pr, nil
}

func (ss *sqlstore) NewProjectID() string {
	return uuid.NewString()
}

func NewSubprojectSqlStore(db *sql.DB, tablename string) store {
	return &sqlstore{db, tablename}
}
