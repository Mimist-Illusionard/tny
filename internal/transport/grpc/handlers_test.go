package grpc

import (
	"context"
	"errors"
	"testing"

	urlgrpc "github.com/Mimist-Illusionard/url-shortener/api/v1/grpc"
	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/Mimist-Illusionard/url-shortener/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type shortenerStub struct {
	create func(context.Context, string) (*domain.URL, error)
	get    func(context.Context, string) (string, error)
}

func (s shortenerStub) CreateShortLink(ctx context.Context, originalURL string) (*domain.URL, error) {
	return s.create(ctx, originalURL)
}

func (s shortenerStub) GetOriginalLink(ctx context.Context, short string) (string, error) {
	return s.get(ctx, short)
}

func TestHandler_CreateShortLink(t *testing.T) {
	t.Parallel()

	stub := shortenerStub{
		create: func(_ context.Context, originalURL string) (*domain.URL, error) {
			return &domain.URL{ID: "id-1", Short: "Abcdef123_", Original: originalURL}, nil
		},
		get: unusedGet,
	}

	response, err := NewHandler(stub).CreateShortLink(context.Background(), &urlgrpc.CreateShortLinkRequest{
		OriginalUrl: "https://example.com",
	})

	require.NoError(t, err)
	assert.Equal(t, "id-1", response.GetId())
	assert.Equal(t, "Abcdef123_", response.GetShortCode())
	assert.Equal(t, "https://example.com", response.GetOriginalUrl())
}

func TestHandler_GetOriginalLink(t *testing.T) {
	t.Parallel()

	stub := shortenerStub{
		create: unusedCreate,
		get:    func(context.Context, string) (string, error) { return "https://example.com", nil },
	}

	response, err := NewHandler(stub).GetOriginalLink(context.Background(), &urlgrpc.GetOriginalLinkRequest{
		ShortCode: "Abcdef123_",
	})

	require.NoError(t, err)
	assert.Equal(t, "https://example.com", response.GetOriginalUrl())
}

func TestMapError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		code codes.Code
	}{
		{name: "invalid URL", err: service.ErrInvalidURL, code: codes.InvalidArgument},
		{name: "invalid short code", err: service.ErrInvalidShortCode, code: codes.InvalidArgument},
		{name: "not found", err: repository.ErrNotFound, code: codes.NotFound},
		{name: "generation exhausted", err: service.ErrCannotGenerate, code: codes.ResourceExhausted},
		{name: "internal", err: errors.New("database error"), code: codes.Internal},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.code, status.Code(mapError(tt.err)))
		})
	}
}

func unusedCreate(context.Context, string) (*domain.URL, error) {
	return nil, errors.New("unexpected CreateShortLink call")
}

func unusedGet(context.Context, string) (string, error) {
	return "", errors.New("unexpected GetOriginalLink call")
}
