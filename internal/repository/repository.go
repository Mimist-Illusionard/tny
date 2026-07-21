package repository

import (
	"errors"

	"github.com/Mimist-Illusionard/url-shortener/internal/config"
	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository/memory"
)

// Repository is a contract for implementations of database repo
// available methods:
// Create creates url and saves it
// Get returns url by its short code
// Delete removes url from the database
type Repository interface {
	Create(url domain.URL) error
	Get(short string) (domain.URL, error)
	Delete(id string)
}

func New(cfg *config.Config) (Repository, error) {
	switch cfg.DatabaseType {
	case config.Memory:
		return memory.New()
	}

	return nil, errors.New("invalid database type")
}
