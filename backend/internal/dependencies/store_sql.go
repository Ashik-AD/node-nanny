package dependencies

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type sqlstore struct {
	dbd            *sql.DB
	tablename      string
	childtablename string
}

func (ss *sqlstore) SaveDependency(dependency *Dependency) error {
	if dependency == nil {
		return errors.New("empty dependency")
	}

	query := fmt.Sprintf("INSERT INTO %s (pid, pkg_id, installed_ver) VALUES (?, ?, ?)", ss.tablename)
	_, err := ss.dbd.Exec(query, dependency.PID, dependency.PkgID, dependency.InstalledVer)
	if err != nil {
		return errors.New("failed to inserting dependency")
	}
	return nil
}

// 'GetDependency' return dependency with package details
func (ss *sqlstore) GetDependency(pid, pkgID string) (*DependencyDetail, error) {
	t1, t2 := ss.tablename, ss.childtablename
	query := fmt.Sprintf(`
        SELECT %s.id, %s.name, %s.npm_link, %s.home_link, 
        %s.repo_link, %s.size, %s.version, %s.installed_ver FROM %s 
        JOIN %s ON %s.pkg_id = %s.id WHERE %s.pid = ? AND %s.pkg_id = ?
        `, t2, t2, t2, t2, t2, t2, t2, t1, t2, t1, t1, t2, t1, t1)

	row := ss.dbd.QueryRow(query, pid, pkgID)
	if err := row.Err(); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("record not found: %v", err)
		}
		return nil, err
	}

	var pkg DependencyDetail

	if err := row.Scan(
		&pkg.ID, &pkg.Name, &pkg.NpmLink, &pkg.HomeLink,
		&pkg.RepoLink, &pkg.Size, &pkg.Version, &pkg.Installed_ver,
	); err != nil {
		return nil, err
	}

	return &pkg, nil
}

// 'GetDependencies' return dependency list with package details
func (ss *sqlstore) GetDependencies(pid string) (*[]DependencyDetail, error) {
	pid = strings.TrimSpace(pid)
	if pid == "" {
		return nil, errors.New("id is empty")
	}
	t1 := ss.tablename
	t2 := ss.childtablename

	query := fmt.Sprintf(`
        SELECT %s.id, %s.name, %s.npm_link, %s.home_link, 
        %s.repo_link, %s.size, %s.version, %s.installed_ver FROM %s 
        JOIN %s ON %s.pkg_id = %s.id WHERE %s.pid = ?
        `, t2, t2, t2, t2, t2, t2, t2, t1, t2, t1, t1, t2, t1)
	rows, err := ss.dbd.Query(query, pid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no records found with th ID: %s", pid)
		}
		return nil, err
	}
	defer rows.Close()

	var dependencies []DependencyDetail
	for rows.Next() {
		var pkg DependencyDetail
		if err := rows.Scan(
			&pkg.ID, &pkg.Name, &pkg.NpmLink,
			&pkg.HomeLink, &pkg.RepoLink,
			&pkg.Size, &pkg.Version,
			&pkg.Installed_ver,
		); err != nil {
			return nil, err
		}
		dependencies = append(dependencies, pkg)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &dependencies, nil
}

// update dependency
func (ss *sqlstore) UpdateDependency(ctx context.Context, pid, pkgId, version string) error {
	query := fmt.Sprintf("UPDATE %s SET version = ? WHERE pid = ? AND pkg_id = ?", ss.tablename)
	_, err := ss.dbd.ExecContext(ctx, query, pid, pkgId)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no record match with pid: %s & pkg_id: %s", pid, pkgId)
		}
		return err
	}
	return nil
}

// save package
func (ss *sqlstore) SavePackage(ctx context.Context, pkg *PackageDetails) (pkgID string, err error) {
	pkgID = ss.NewID()
	pkg.ID = pkgID

	query := fmt.Sprintf(`
        NSERT INTO %s 
        (id, name, version, npm_link, home_link, repo_link, size)
        VALUES (?, ?, ?, ?, ?, ?, ?)
   `, ss.childtablename)
	_, err = ss.dbd.ExecContext(ctx, query, pkg.ID, pkg.Name, pkg.Version, pkg.NpmLink, pkg.HomeLink, pkg.RepoLink, pkg.Size)
	if err != nil {
		return "", err
	}
	return pkgID, nil
}

// update package details
func (ss *sqlstore) UpdatePackage(ctx context.Context, pkg *PackageDetails) (*PackageDetails, error) {
	query := fmt.Sprintf("UPDATE %s SET name = ?, version = ?, npm_link = ?, home_link = ?, repo_link = ?, size = ?", ss.tablename)
	res, err := ss.dbd.ExecContext(ctx, query, pkg.Name, pkg.Version, pkg.NpmLink, pkg.HomeLink, pkg.RepoLink, pkg.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no record match with the pkg_id: %s", pkg.ID)
		}
		return nil, err
	}
	if _, err := res.RowsAffected(); err != nil {
		return nil, err
	}
	return pkg, nil
}

// remove dependency
func (ss *sqlstore) RemoveDependency(pid, pkgID string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE pid = ? AND pkg_id = ?", ss.tablename)
	_, err := ss.dbd.Exec(query, pid, pkgID)
	if err != nil {
		return err
	}
	return nil
}

func (ss *sqlstore) NewID() string {
	return uuid.NewString()
}

func NewDependencyStore(db *sql.DB, tablename, childtablename string) store {
	return &sqlstore{db, tablename, childtablename}
}
