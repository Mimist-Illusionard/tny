package memory

import (
	"context"
	"sync"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
)

type MemoryRepository struct {
	shortUrls map[string]*domain.URL
	origUrls  map[string]string
	mu        sync.RWMutex
}

func New() (*MemoryRepository, error) {
	return &MemoryRepository{
		shortUrls: make(map[string]*domain.URL),
		origUrls:  make(map[string]string)}, nil
}

func (r *MemoryRepository) Create(ctx context.Context, url *domain.URL) error {
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

func (r *MemoryRepository) Get(ctx context.Context, short string) (*domain.URL, error) {
	r.mu.RLock()
	url, ok := r.shortUrls[short]
	r.mu.RUnlock()
	if !ok {
		return nil, repository.ErrNotFound
	}
	return url, nil
}

func (r *MemoryRepository) GetByOriginalURL(ctx context.Context, url string) (*domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	short, ok := r.origUrls[url]
	if !ok {
		return nil, repository.ErrNotFound
	}

	return r.shortUrls[short], nil
}

func (r *MemoryRepository) Delete(ctx context.Context, short string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u, ok := r.shortUrls[short]
	if !ok {
		return
	}
	delete(r.shortUrls, short)
	delete(r.origUrls, u.Original)
}
