package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Database string

type DatabaseParams struct {
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
}

const (
	Postgres Database = "postgres"
	Memory   Database = "memory"
)

type Config struct {
	DBType   Database
	DBParams DatabaseParams
	HTTPPort string
	GRPCPort string
}

func New(db, httpPort, grpcPort, envPath string) (*Config, error) {
	if envPath != "" {
		if err := godotenv.Load(envPath); err != nil {
			return nil, fmt.Errorf("load env file %q: %w", envPath, err)
		}
	}

	params := DatabaseParams{}
	if err := env.Parse(&params); err != nil {
		return nil, fmt.Errorf("parse database environment: %w", err)
	}

	cfg := &Config{
		DBParams: params,
		HTTPPort: httpPort,
		GRPCPort: grpcPort,
	}

	switch Database(db) {
	case Postgres:
		cfg.DBType = Postgres
		if err := validatePostgresParams(params); err != nil {
			return nil, err
		}
	case Memory:
		cfg.DBType = Memory
	default:
		return nil, fmt.Errorf("unsupported database %q", db)
	}

	return cfg, nil
}

func validatePostgresParams(params DatabaseParams) error {
	missing := make([]string, 0, 5)

	if params.Name == "" {
		missing = append(missing, "DB_NAME")
	}
	if params.Host == "" {
		missing = append(missing, "DB_HOST")
	}
	if params.Port == "" {
		missing = append(missing, "DB_PORT")
	}
	if params.Username == "" {
		missing = append(missing, "DB_USERNAME")
	}
	if params.Password == "" {
		missing = append(missing, "DB_PASSWORD")
	}

	if len(missing) == 0 {
		return nil
	}

	return fmt.Errorf("missing PostgreSQL environment variables: %s", strings.Join(missing, ", "))
}
