package domain

import (
	"time"

	"github.com/google/uuid"
)

type URL struct {
	ID        string    `json:"id"`
	Short     string    `json:"shortUrl"`
	Original  string    `json:"originalUrl"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// NewURL creates url for future use
func NewURL(origURL, shortURL string) *URL {
	return &URL{
		ID:        uuid.New().String(),
		Original:  origURL,
		Short:     shortURL,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour),
	}
}
