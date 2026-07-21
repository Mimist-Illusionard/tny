package config

type Database string

const (
	Postgres Database = "postgres"
	Memory   Database = "memory"
)

type Config struct {
	DatabaseType Database
	Port         string
}

func New(db, port string) *Config {
	cfg := &Config{Port: port}

	switch db {
	case "postgres":
		cfg.DatabaseType = Postgres
	case "memory":
		cfg.DatabaseType = Memory
	}

	return cfg
}
