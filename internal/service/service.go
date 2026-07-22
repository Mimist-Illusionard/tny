package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
)

var (
	ErrCannotGenerate = errors.New("can't generate short url")
)

const (
	maxAttempts = 5
	alphabet    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	codeLength  = 10
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateShortLink(ctx context.Context, url string) (*domain.URL, error) {
	for i := 0; i < maxAttempts; i++ {
		short, err := generateShort(codeLength)
		if err != nil {
			return nil, err
		}

		u := domain.NewUrl(url, short)
		err = s.repo.Create(ctx, u)
		if errors.Is(err, repository.ErrExists) {
			continue
		}

		if errors.Is(err, repository.ErrNotUnique) {
			return s.repo.GetByOriginal(ctx, url)
		}

		if err != nil {
			return nil, err
		}

		return u, nil
	}

	return nil, ErrCannotGenerate
}

func (s *Service) GetOriginalLink(ctx context.Context, short string) (string, error) {
	u, err := s.repo.Get(ctx, short)
	if errors.Is(err, repository.ErrNotFound) {
		return "", err
	}

	if err != nil {
		return "", err
	}

	return u.Original, err
}

func generateShort(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be positive")
	}

	code := make([]byte, length)
	limit := big.NewInt(int64(len(alphabet)))

	for i := range code {
		index, err := rand.Int(rand.Reader, limit)
		if err != nil {
			return "", fmt.Errorf("generate short code: %w", err)
		}

		code[i] = alphabet[index.Int64()]
	}

	return string(code), nil
}
