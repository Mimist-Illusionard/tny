package app

import (
	"errors"

	"github.com/Mimist-Illusionard/url-shortener/internal/config"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository/memory"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository/postgres"
	"github.com/Mimist-Illusionard/url-shortener/internal/transport/http"
)

func Run(cfg *config.Config) error {
	repo, err := newRepository(cfg)
	if err != nil {
		return err
	}

	return http.Serve(cfg, repo)
}

func newRepository(cfg *config.Config) (repository.Repository, error) {
	switch cfg.DBType {
	case config.Memory:
		return memory.New(), nil
	case config.Postgres:
		sql, err := postgres.Connect(cfg)
		if err != nil {
			return nil, err
		}
		return postgres.New(sql), nil
	}

	return nil, errors.New("invalid database type")
}
