package projects

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

func (ss *sqlstore) SaveProject(ctx context.Context, project *Project) (string, error) {
	project.ID = ss.NewProjectID()

	query := fmt.Sprintf(`
        INSERT INTO %s (id, name, path, detail, files, 
        dependencies, subprojects, logs) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		ss.tablename)

	stm, err := ss.dbd.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stm.Close()

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

func (ss *sqlstore) GetProjects() (*[]Project, error) {
	query := fmt.Sprintf("SELECT * FROM %s", ss.tablename)
	stm, err := ss.dbd.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	rows, err := stm.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var pr Project
		if err := rows.Scan(
			&pr.ID, &pr.Name, &pr.Path,
			&pr.Detail, &pr.Files,
			&pr.Dependencies, &pr.Subproject,
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

func (ss *sqlstore) GetProjectByID(pid string) (*Project, error) {
	var pr Project

	pid = strings.TrimSpace(pid)
	if pid == "" {
		return nil, errors.New("empty project ID")
	}

	row := ss.dbd.QueryRow("SELECT * FROM ? WHERE id = ?", ss.tablename, pid)
	if err := row.Scan(
		&pr.ID, &pr.Name, &pr.Path,
		&pr.Detail, &pr.Files, &pr.Dependencies,
		&pr.Subproject, &pr.Logs,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No record found with the ID: %s", pid)
		}
		return nil, err
	}

	return &pr, nil
}

func (ss *sqlstore) GetDetailedProject(pid string) (*Project, error) {
	// implement this method
	var pr Project

	pid = strings.TrimSpace(pid)
	if pid == "" {
		return nil, errors.New("empty project ID")
	}
	return &pr, nil
}

func (ss *sqlstore) NewProjectID() string {
	return uuid.NewString()
}

func NewProjectSqlStore(db *sql.DB, tablename string) store {
	return &sqlstore{db, tablename}
}
