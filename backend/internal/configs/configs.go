package configs

type Configs struct{}

func (cfg *Configs) ProjectSqlTable() string {
	return "projects"
}

func (cfg *Configs) DependencySqlTable() string {
	return "dependencies"
}

func (cfg *Configs) PackgeDetailsSqlTable() string {
	return "pkgdetails"
}

func (cfg *Configs) SubprojectSqlTable() string {
	return "subprojects"
}

func (cfg *Configs) FilesSqlTable() string {
	return "files"
}
