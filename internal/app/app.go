package app

import (
	"errors"

	"github.com/Mimist-Illusionard/url-shortener/internal/config"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository/memory"
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
	switch cfg.DatabaseType {
	case config.Memory:
		return memory.New()
	}

	return nil, errors.New("invalid database type")
}
