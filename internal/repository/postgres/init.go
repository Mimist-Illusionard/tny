package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Mimist-Illusionard/tny/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBParams.Host, cfg.DBParams.Username, cfg.DBParams.Password, cfg.DBParams.Name, cfg.DBParams.Port,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	migrationDSN := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBParams.Username,
		cfg.DBParams.Password,
		cfg.DBParams.Host,
		cfg.DBParams.Port,
		cfg.DBParams.Name,
	)

	if err := RunMigration(migrationDSN); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
