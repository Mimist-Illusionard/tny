package config

type Database string

type DatabaseParams struct {
	Name     string
	Host     string
	Port     string
	Username string
	Password string
}

const (
	Postgres Database = "postgres"
	Memory   Database = "memory"
)

type Config struct {
	DBType   Database
	DBParams DatabaseParams
	Port     string
}

func New(db, port string) *Config {
	cfg := &Config{Port: port}

	switch db {
	case "postgres":
		cfg.DBType = Postgres
	case "memory":
		cfg.DBType = Memory
	}

	return cfg
}
