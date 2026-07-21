package memory

import (
	"errors"
	"sync"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
)

type MemRepository struct {
	urls map[string]*domain.URL
	mu   sync.RWMutex
}

func New() (*MemRepository, error) {
	return &MemRepository{}, nil
}

func (r *MemRepository) Create(url domain.URL) error {
	r.mu.RLock()
	if _, ok := r.urls[url.Short]; ok {
		return errors.New("url already exists")
	}
	r.mu.RUnlock()

	r.mu.Lock()
	r.urls[url.ID] = &url
	r.mu.Unlock()
	return nil
}

func (r *MemRepository) Get(short string) (domain.URL, error) {
	r.mu.RLock()
	url, ok := r.urls[short]
	r.mu.RUnlock()
	if !ok {
		return domain.URL{}, errors.New("url not found")
	}
	return *url, nil
}

func (r *MemRepository) Delete(id string) {
	r.mu.Lock()
	delete(r.urls, id)
	r.mu.Unlock()
}
