package config

import (
	"log"

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
	Port     string
}

func New(db, port, envPath string) *Config {
	cfg := &Config{Port: port}

	if envPath != "" {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		params := DatabaseParams{}
		err = env.Parse(&params)
		if err != nil {
			log.Fatalf("Error parsing .env file")
		}

		cfg.DBParams = params
		log.Println(cfg.DBParams)
	}

	switch db {
	case "postgres":
		cfg.DBType = Postgres
	case "memory":
		cfg.DBType = Memory
	}

	return cfg
}
