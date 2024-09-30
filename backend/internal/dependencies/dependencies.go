package dependencies

import (
	"context"
	"errors"
	"strings"
)

type Dependency struct {
	PID          string
	PkgID        string
	InstalledVer string
}

type PackageDetails struct {
	ID       string
	Name     string
	Version  string
	NpmLink  string
	HomeLink string
	RepoLink string
	Size     float32
}

type store interface {
	SaveDependency(*Dependency) error
	SavePackage(ctx context.Context, pkg *PackageDetails) (pkgId string, err error)
	GetDependency(pid string, pkgID string) (*DependencyDetail, error)
	GetDependencies(pid string) (*[]DependencyDetail, error)
	UpdateDependency(ctx context.Context, pid string, pkgID string, version string) error
	UpdatePackage(ctx context.Context, pkg *PackageDetails) (*PackageDetails, error)
	RemoveDependency(pid string, pkgID string) error
}

type DependencyDetail struct {
	Installed_ver string
	PackageDetails
}

type Dependencies struct {
	store store
}

func (d *Dependency) Sanitize() {
	d.PID = strings.TrimSpace(d.PID)
	d.PkgID = strings.TrimSpace(d.PkgID)
	d.InstalledVer = strings.TrimSpace(d.InstalledVer)
}

func (pkg *PackageDetails) Sanitize() {
	pkg.ID = trimSpace(pkg.ID)
	pkg.Name = trimSpace(pkg.Name)
	pkg.Version = trimSpace(pkg.Version)
	pkg.NpmLink = trimSpace(pkg.NpmLink)
	pkg.HomeLink = trimSpace(pkg.HomeLink)
	pkg.RepoLink = trimSpace(pkg.RepoLink)
}

func (d *Dependency) ValidateForCreate() error {
	if d == nil {
		return errors.New("empty dependency")
	}

	d.Sanitize()
	if d.PID == "" {
		return errors.New("dependency pid can't be empty")
	}
	if d.PkgID == "" {
		return errors.New("dependency pkgid can't be empty")
	}
	if d.InstalledVer == "" {
		return errors.New("dependency version to be install can't be empty")
	}
	return nil
}

func (pkg *PackageDetails) ValidateForCreate() error {
	if pkg == nil {
		return errors.New("empty package details")
	}

	pkg.Sanitize()
	if pkg.Name == "" {
		return errors.New("packge name can't be empty")
	}
	if pkg.Version == "" {
		return errors.New("packge version can't be empty")
	}
	if pkg.Size <= 0 {
		return errors.New("packge size can't be 0")
	}
	if pkg.NpmLink == "" {
		return errors.New("packge npm link can't be empty")
	}
	if pkg.RepoLink == "" {
		return errors.New("packge repository link can't be empty")
	}
	return nil
}

func (dp *Dependencies) SaveDependency(dependency *Dependency) error {
	err := dependency.ValidateForCreate()
	if err != nil {
		return nil
	}

	err = dp.store.SaveDependency(dependency)
	if err != nil {
		return nil
	}
	return nil
}

func (dp *Dependencies) SavePackage(ctx context.Context, pkg *PackageDetails) (pkgID string, err error) {
	err = pkg.ValidateForCreate()
	if err != nil {
		return "", err
	}
	pkg.ID, err = dp.store.SavePackage(ctx, pkg)
	if err != nil {
		return "", err
	}
	return pkg.ID, nil
}

func (dp *Dependencies) GetDependency(pid string, pkgID string) (*DependencyDetail, error) {
	if trimSpace(pid) == "" || trimSpace(pkgID) == "" {
		return nil, errors.New("projec id or package id is empty")
	}
	dpDetails, err := dp.store.GetDependency(pid, pkgID)
	if err != nil {
		return nil, err
	}
	return dpDetails, nil
}

func (dp *Dependencies) GetDependencies(pid string) (*[]DependencyDetail, error) {
	if trimSpace(pid) == "" {
		return nil, errors.New("dependency pid can't be empty")
	}
	dpLists, err := dp.store.GetDependencies(pid)
	if err != nil {
		return nil, err
	}
	return dpLists, nil
}

func (dp *Dependencies) UpdateDependency(ctx context.Context, pid string, pkgID string, version string) error {
	if trimSpace(pid) == "" || trimSpace(pkgID) == "" {
		return errors.New("projec id or package id can't be empty")
	}

	if trimSpace(version) == "" {
		return errors.New("version is empty")
	}

	err := dp.store.UpdateDependency(ctx, pid, pkgID, version)
	if err != nil {
		return err
	}
	return nil
}

func (dp *Dependencies) UpdatePackage(ctx context.Context, pkg *PackageDetails) (*PackageDetails, error) {
	err := pkg.ValidateForCreate()
	if err != nil {
		return nil, err
	}
	pkg, err = dp.store.UpdatePackage(ctx, pkg)
	if err != nil {
		return nil, err
	}
	return pkg, nil
}

func (dp *Dependencies) RemoveDependency(pid string, pkgID string) error {
	if trimSpace(pid) == "" || trimSpace(pkgID) == "" {
		return errors.New("pid or pkgid is empty")
	}
	err := dp.store.RemoveDependency(pid, pkgID)
	if err != nil {
		return err
	}
	return nil
}

func NewDependencySvc(store store) *Dependencies {
	return &Dependencies{store: store}
}


func trimSpace(val string) string {
	return strings.TrimSpace(val)
}
