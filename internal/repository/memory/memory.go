package memory

import (
	"context"
	"sync"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
)

type MemRepository struct {
	shortUrls map[string]*domain.URL
	origUrls  map[string]string
	mu        sync.RWMutex
}

func New() (*MemRepository, error) {
	return &MemRepository{
		shortUrls: make(map[string]*domain.URL),
		origUrls:  make(map[string]string)}, nil
}

func (r *MemRepository) Create(ctx context.Context, url *domain.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.shortUrls[url.Short]; ok {
		return repository.ErrExists
	}

	if _, ok := r.origUrls[url.Original]; ok {
		return repository.ErrNotUnique
	}

	r.shortUrls[url.Short] = url
	r.origUrls[url.Original] = url.Short

	return nil
}

func (r *MemRepository) Get(ctx context.Context, short string) (*domain.URL, error) {
	r.mu.RLock()
	url, ok := r.shortUrls[short]
	r.mu.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return url, nil
}

func (r *MemRepository) GetByOriginalURL(ctx context.Context, url string) (*domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	short, ok := r.origUrls[url]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return r.shortUrls[short], nil
}

func (r *MemRepository) Delete(ctx context.Context, short string) {
	r.mu.Lock()

	u := r.shortUrls[short]
	delete(r.shortUrls, short)
	delete(r.origUrls, u.Original)

	r.mu.Unlock()
}
