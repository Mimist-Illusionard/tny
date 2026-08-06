package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Mimist-Illusionard/url-shortener/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBParams.Host, cfg.DBParams.Username, cfg.DBParams.Password, cfg.DBParams.Name, cfg.DBParams.Port,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = RunMigration(dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
