package repository

import (
	"context"
	"errors"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
)

var (
	ErrNotFound  = errors.New("url not found")
	ErrExists    = errors.New("short url already exists")
	ErrNotUnique = errors.New("link to url already exists")
)

// Repository is a contract for implementations of database repo
// available methods:
// Create creates url and saves it
// Get returns url by its short code
// Delete removes url from the database
//
//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=repository Repository
type Repository interface {
	Create(ctx context.Context, url *domain.URL) error
	Get(ctx context.Context, short string) (*domain.URL, error)
	GetByOriginalURL(ctx context.Context, url string) (*domain.URL, error)
	Delete(ctx context.Context, short string)
}
