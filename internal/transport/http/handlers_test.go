package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/Mimist-Illusionard/url-shortener/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestCreateShortLink(t *testing.T) {
	t.Parallel()

	t.Run("returns created link", func(t *testing.T) {
		t.Parallel()

		stub := shortenerStub{
			create: func(_ context.Context, originalURL string) (*domain.URL, error) {
				assert.Equal(t, "https://example.com", originalURL)
				return &domain.URL{ID: "id-1", Short: "Abcdef123_", Original: originalURL}, nil
			},
			get: unusedGet,
		}

		req := httptest.NewRequest(http.MethodPost, "/api/v1/links", strings.NewReader(`{"url":"https://example.com"}`))
		recorder := httptest.NewRecorder()

		RegisterHandlers(stub).ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code)
		assert.Contains(t, recorder.Header().Get("Content-Type"), "application/json")
		assert.JSONEq(t, `{"id":"id-1","shortCode":"Abcdef123_","originalUrl":"https://example.com"}`, recorder.Body.String())
	})

	t.Run("rejects malformed JSON", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodPost, "/api/v1/links", strings.NewReader(`{"url":`))
		recorder := httptest.NewRecorder()

		RegisterHandlers(shortenerStub{create: unusedCreate, get: unusedGet}).ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("maps invalid URL to 400", func(t *testing.T) {
		t.Parallel()

		stub := shortenerStub{
			create: func(context.Context, string) (*domain.URL, error) { return nil, service.ErrInvalidURL },
			get:    unusedGet,
		}
		req := httptest.NewRequest(http.MethodPost, "/api/v1/links", strings.NewReader(`{"url":"bad"}`))
		recorder := httptest.NewRecorder()

		RegisterHandlers(stub).ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
}

func TestGetOriginalLink(t *testing.T) {
	t.Parallel()

	const short = "Abcdef123_"

	t.Run("returns original URL as JSON", func(t *testing.T) {
		t.Parallel()

		stub := shortenerStub{
			create: unusedCreate,
			get: func(_ context.Context, gotShort string) (string, error) {
				assert.Equal(t, short, gotShort)
				return "https://example.com", nil
			},
		}
		req := httptest.NewRequest(http.MethodGet, "/api/v1/links/"+short, nil)
		recorder := httptest.NewRecorder()

		RegisterHandlers(stub).ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.JSONEq(t, `{"originalUrl":"https://example.com"}`, recorder.Body.String())
	})

	t.Run("maps not found to 404", func(t *testing.T) {
		t.Parallel()

		stub := shortenerStub{
			create: unusedCreate,
			get:    func(context.Context, string) (string, error) { return "", repository.ErrNotFound },
		}
		req := httptest.NewRequest(http.MethodGet, "/api/v1/links/"+short, nil)
		recorder := httptest.NewRecorder()

		RegisterHandlers(stub).ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func TestRedirectToOriginalLink(t *testing.T) {
	t.Parallel()

	stub := shortenerStub{
		create: unusedCreate,
		get:    func(context.Context, string) (string, error) { return "https://example.com", nil },
	}
	req := httptest.NewRequest(http.MethodGet, "/Abcdef123_", nil)
	recorder := httptest.NewRecorder()

	RegisterHandlers(stub).ServeHTTP(recorder, req)

	response := recorder.Result()
	require.NoError(t, response.Body.Close())
	assert.Equal(t, http.StatusFound, response.StatusCode)
	assert.Equal(t, "https://example.com", response.Header.Get("Location"))
}

func unusedCreate(context.Context, string) (*domain.URL, error) {
	return nil, errors.New("unexpected CreateShortLink call")
}

func unusedGet(context.Context, string) (string, error) {
	return "", errors.New("unexpected GetOriginalLink call")
}
