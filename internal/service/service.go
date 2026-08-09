package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strings"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
)

var (
	ErrCannotGenerate   = errors.New("can't generate short url")
	ErrInvalidURL       = errors.New("invalid original url")
	ErrInvalidShortCode = errors.New("invalid short code")
)

const (
	maxAttempts          = 5
	alphabet             = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	codeLength           = 10
	maxOriginalURLLength = 2048
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateShortLink(ctx context.Context, originalURL string) (*domain.URL, error) {
	if !isValidOriginalURL(originalURL) {
		return nil, ErrInvalidURL
	}

	for i := 0; i < maxAttempts; i++ {
		short, err := generateShort(codeLength)
		if err != nil {
			return nil, err
		}

		u := domain.NewURL(originalURL, short)
		err = s.repo.Create(ctx, u)
		if errors.Is(err, repository.ErrExists) {
			continue
		}

		if errors.Is(err, repository.ErrNotUnique) {
			return s.repo.GetByOriginalURL(ctx, originalURL)
		}

		if err != nil {
			return nil, err
		}

		return u, nil
	}

	return nil, ErrCannotGenerate
}

func (s *Service) GetOriginalLink(ctx context.Context, short string) (string, error) {
	if !isValidShortCode(short) {
		return "", ErrInvalidShortCode
	}

	u, err := s.repo.Get(ctx, short)
	if err != nil {
		return "", err
	}

	return u.Original, nil
}

func generateShort(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be positive")
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

func isValidOriginalURL(value string) bool {
	if strings.TrimSpace(value) != value || value == "" || len(value) > maxOriginalURLLength || strings.ContainsAny(value, " \t\r\n") {
		return false
	}

	parsed, err := url.ParseRequestURI(value)
	if err != nil {
		return false
	}

	return (parsed.Scheme == "http" || parsed.Scheme == "https") && parsed.Host != ""
}

func isValidShortCode(short string) bool {
	if len(short) != codeLength {
		return false
	}

	for i := range short {
		if !strings.ContainsRune(alphabet, rune(short[i])) {
			return false
		}
	}

	return true
}
