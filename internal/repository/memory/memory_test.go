package memory

import (
	"context"
	"testing"

	"github.com/Mimist-Illusionard/tny/internal/domain"
	"github.com/Mimist-Illusionard/tny/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("creates URL", func(t *testing.T) {
		t.Parallel()

		r := New()

		want := domain.NewURL(
			"https://example.com/page",
			"abc123",
		)

		err = r.Create(ctx, want)
		require.NoError(t, err)

		gotByShort, err := r.Get(ctx, want.Short)
		require.NoError(t, err)

		assertURL(t, want, gotByShort)

		gotByOriginal, err := r.GetByOriginalURL(ctx, want.Original)
		require.NoError(t, err)

		assertURL(t, want, gotByOriginal)
	})

	t.Run("returns ErrExists when the same URL already exists", func(t *testing.T) {
		t.Parallel()

		r := New()

		url := domain.NewURL(
			"https://example.com/page",
			"abc123",
		)

		require.NoError(t, r.Create(ctx, url))

		err = r.Create(ctx, url)

		require.ErrorIs(t, err, repository.ErrExists)

		got, getErr := r.Get(ctx, url.Short)
		require.NoError(t, getErr)

		assertURL(t, url, got)
	})

	t.Run("returns ErrNotUnique when original URL already exists", func(t *testing.T) {
		t.Parallel()

		r := New()

		existingURL := domain.NewURL(
			"https://example.com/page",
			"abc123",
		)

		newURL := domain.NewURL(
			existingURL.Original,
			"different",
		)

		require.NoError(t, r.Create(ctx, existingURL))

		err = r.Create(ctx, newURL)

		require.ErrorIs(t, err, repository.ErrNotUnique)

		got, getErr := r.GetByOriginalURL(ctx, existingURL.Original)
		require.NoError(t, getErr)

		assertURL(t, existingURL, got)

		_, getErr = r.Get(ctx, newURL.Short)
		require.ErrorIs(t, getErr, repository.ErrNotFound)
	})
}

func TestRepository_Get(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name      string
		shortCode string
		prepare   bool
		wantErr   error
	}{
		{
			name:      "returns existing URL",
			shortCode: "abc123",
			prepare:   true,
		},
		{
			name:      "returns ErrNotFound for missing URL",
			shortCode: "missing",
			wantErr:   repository.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := New()

			want := domain.NewURL(
				"https://example.com/page",
				"abc123",
			)

			if tt.prepare {
				require.NoError(t, r.Create(ctx, want))
			}

			got, err := r.Get(ctx, tt.shortCode)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)

			assertURL(t, want, got)
		})
	}
}

func TestRepository_GetByOriginalURL(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const existingOriginalURL = "https://example.com/page"

	tests := []struct {
		name        string
		originalURL string
		prepare     bool
		wantErr     error
	}{
		{
			name:        "returns existing URL",
			originalURL: existingOriginalURL,
			prepare:     true,
		},
		{
			name:        "returns ErrNotFound for missing original URL",
			originalURL: "https://missing.example.com/page",
			wantErr:     repository.ErrNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := New()

			want := domain.NewURL(
				existingOriginalURL,
				"abc123",
			)

			if tt.prepare {
				require.NoError(t, r.Create(ctx, want))
			}

			got, err := r.GetByOriginalURL(ctx, tt.originalURL)

			if tt.wantErr != nil {
				require.ErrorIs(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)

			assertURL(t, want, got)
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("deletes URL from both indexes", func(t *testing.T) {
		t.Parallel()

		r := New()

		url := domain.NewURL(
			"https://example.com/page",
			"abc123",
		)

		require.NoError(t, r.Create(ctx, url))

		r.Delete(ctx, url.Short)

		_, err = r.Get(ctx, url.Short)
		require.ErrorIs(t, err, repository.ErrNotFound)

		_, err = r.GetByOriginalURL(ctx, url.Original)
		require.ErrorIs(t, err, repository.ErrNotFound)
	})

	t.Run("allows creating URL again after deletion", func(t *testing.T) {
		t.Parallel()

		r := New()

		url := domain.NewURL(
			"https://example.com/page",
			"abc123",
		)

		require.NoError(t, r.Create(ctx, url))

		r.Delete(ctx, url.Short)

		err = r.Create(ctx, url)
		require.NoError(t, err)

		got, err := r.Get(ctx, url.Short)
		require.NoError(t, err)

		assertURL(t, url, got)
	})

	t.Run("does nothing when URL does not exist", func(t *testing.T) {
		t.Parallel()

		r := New()

		existingURL := domain.NewURL(
			"https://example.com/page",
			"abc123",
		)

		require.NoError(t, r.Create(ctx, existingURL))

		r.Delete(ctx, "missing")

		got, err := r.Get(ctx, existingURL.Short)
		require.NoError(t, err)

		assertURL(t, existingURL, got)
	})
}

func assertURL(t *testing.T, want, got *domain.URL) {
	t.Helper()

	require.NotNil(t, want)
	require.NotNil(t, got)

	assert.Equal(t, want.Original, got.Original)
	assert.Equal(t, want.Short, got.Short)
}
