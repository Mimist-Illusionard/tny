package service

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/Mimist-Illusionard/tny/internal/domain"
	"github.com/Mimist-Illusionard/tny/internal/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_CreateShortLink(t *testing.T) {
	t.Parallel()

	const originalURL = "https://example.com/page"

	tests := []struct {
		name      string
		prepare   func(*repository.MockRepository)
		want      *domain.URL
		wantError error
	}{
		{
			name: "creates a short link",
			prepare: func(r *repository.MockRepository) {
				r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name: "retries when generated short code collides",
			prepare: func(r *repository.MockRepository) {
				gomock.InOrder(
					r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(repository.ErrExists),
					r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil),
				)
			},
		},
		{
			name: "returns existing link for the same original URL",
			prepare: func(r *repository.MockRepository) {
				existing := &domain.URL{Original: originalURL, Short: "Abcdef123_"}
				r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(repository.ErrNotUnique)
				r.EXPECT().GetByOriginalURL(gomock.Any(), originalURL).Return(existing, nil)
			},
			want: &domain.URL{Original: originalURL, Short: "Abcdef123_"},
		},
		{
			name: "returns repository error",
			prepare: func(r *repository.MockRepository) {
				r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
			},
			wantError: errors.New("database error"),
		},
		{
			name: "returns ErrCannotGenerate after all collisions",
			prepare: func(r *repository.MockRepository) {
				r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(repository.ErrExists).Times(maxAttempts)
			},
			wantError: ErrCannotGenerate,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			r := repository.NewMockRepository(ctrl)
			tt.prepare(r)

			s := NewService(r)
			got, err := s.CreateShortLink(context.Background(), originalURL)

			if tt.wantError != nil {
				require.Error(t, err)
				assert.Equal(t, tt.wantError.Error(), err.Error())
				assert.Nil(t, got)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, got)
			assert.Equal(t, originalURL, got.Original)
			assert.Len(t, got.Short, codeLength)
			if tt.want != nil {
				assert.Equal(t, tt.want.Short, got.Short)
			}
		})
	}
}

func TestService_CreateShortLink_RejectsInvalidURL(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	r := repository.NewMockRepository(ctrl)
	s := NewService(r)

	for _, value := range []string{"", "example.com", "ftp://example.com", " https://example.com", "https://example.com/a b", "https://example.com/" + strings.Repeat("a", maxOriginalURLLength)} {
		value := value
		t.Run(value, func(t *testing.T) {
			t.Parallel()

			got, err := s.CreateShortLink(context.Background(), value)
			require.ErrorIs(t, err, ErrInvalidURL)
			assert.Nil(t, got)
		})
	}
}

func TestService_GetOriginalLink(t *testing.T) {
	t.Parallel()

	const (
		short       = "Abcdef123_"
		originalURL = "https://example.com/page"
	)

	t.Run("returns original URL", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		r := repository.NewMockRepository(ctrl)
		r.EXPECT().Get(gomock.Any(), short).Return(&domain.URL{Original: originalURL, Short: short}, nil)

		got, err := NewService(r).GetOriginalLink(context.Background(), short)
		require.NoError(t, err)
		assert.Equal(t, originalURL, got)
	})

	t.Run("returns ErrNotFound", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		r := repository.NewMockRepository(ctrl)
		r.EXPECT().Get(gomock.Any(), short).Return(nil, repository.ErrNotFound)

		got, err := NewService(r).GetOriginalLink(context.Background(), short)
		require.ErrorIs(t, err, repository.ErrNotFound)
		assert.Empty(t, got)
	})

	t.Run("rejects invalid short code", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		r := repository.NewMockRepository(ctrl)

		got, err := NewService(r).GetOriginalLink(context.Background(), "too-short")
		require.ErrorIs(t, err, ErrInvalidShortCode)
		assert.Empty(t, got)
	})
}

func TestGenerateShort(t *testing.T) {
	t.Parallel()

	t.Run("uses requested length and alphabet", func(t *testing.T) {
		t.Parallel()

		result, err := generateShort(codeLength)
		require.NoError(t, err)
		assert.Len(t, result, codeLength)

		for _, char := range result {
			assert.Contains(t, alphabet, string(char))
		}
	})

	t.Run("rejects non-positive length", func(t *testing.T) {
		t.Parallel()

		result, err := generateShort(0)
		require.Error(t, err)
		assert.Empty(t, result)
	})
}
