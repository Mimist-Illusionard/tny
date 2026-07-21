package config

type Database string

const (
	Postgres Database = "postgres"
	Memory   Database = "memory"
)

type Config struct {
	DatabaseType Database
}

func New(db string) *Config {
	switch db {
	case "postgres":
		return &Config{Postgres}
	case "memory":
		return &Config{Memory}
	}

	return &Config{Memory}
}
