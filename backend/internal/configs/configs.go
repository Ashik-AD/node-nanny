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
