package files

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type sqlstore struct {
	dbd       *sql.DB
	tablename string
}

func (ss *sqlstore) SaveFiles(files *[]File) (*[]File, error) {
	query := fmt.Sprintf("INSERT INTO %s (id, pid, name, path) VALUES (?, ?, ?, ?)", ss.tablename)
	stm, err := ss.dbd.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	for _, file := range *files {
		file.ID = ss.NewFileID()
		_, err = stm.Exec(file.ID, file.PID, file.Name, file.Path)
		if err != nil {
			return nil, err
		}
	}

	return files, nil
}

func (ss *sqlstore) GetFilesByProject(pid string) (*[]File, error) {
	var files []File

	query := fmt.Sprintf("SELECT * FROM %s WHERE pid = ?", ss.tablename)
	rc, err := ss.dbd.Query(query, pid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no records match with the pid: %s", pid)
		}
		return nil, err
	}

	for rc.Next() {
		var file File
		if err := rc.Scan(&file.ID, &file.PID, &file.Name, &file.Path); err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	if rc.Err() != nil {
		return nil, err
	}
	return &files, nil
}

func (ss *sqlstore) GetFileByID(id string) (*File, error) {
	var file File
	var err error

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", ss.tablename)
	row := ss.dbd.QueryRow(query, id)
	err = row.Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no file found with ID: %s", id)
		}
		return nil, err
	}
	if err = row.Scan(&file.ID, &file.PID, &file.Name, &file.Path); err != nil {
		return nil, err
	}
	return &file, nil
}

func (ss *sqlstore) NewFileID() string {
	return uuid.NewString()
}

func NewFileSqlStore(db *sql.DB, tablename string) store {
	return &sqlstore{
		dbd:       db,
		tablename: tablename,
	}
}
