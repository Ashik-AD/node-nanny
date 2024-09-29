package configs

type Configs struct{}

func (cfg *Configs) ProjectSqlTable() string {
	return "projects"
}
