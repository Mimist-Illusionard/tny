package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateShortLink_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := repository.NewMockRepository(ctrl)
	s := NewService(r)
	originalURL := "https://example.com"
	ctx := context.Background()

	r.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	u, err := s.CreateShortLink(ctx, originalURL)
	require.NoError(t, err)
	assert.Equal(t, originalURL, u)
	assert.NotEmpty(t, u.Short)
	assert.Len(t, u.Short, codeLength)
}

func TestCreateShortLink_ErrExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := repository.NewMockRepository(ctrl)
	s := NewService(r)
	URL := "https://example.com"
	ctx := context.Background()

	r.EXPECT().Create(ctx, gomock.Any()).Return(repository.ErrExists)
	r.EXPECT().Create(ctx, gomock.Any()).Return(nil)

	u, err := s.CreateShortLink(ctx, URL)

	require.Error(t, err)
	assert.Equal(t, URL, u.Original)
	assert.NotEmpty(t, u.Short)
}

func TestCreateShortLink_ErrNotUnique(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := repository.NewMockRepository(ctrl)
	s := NewService(r)
	URL := "https://example.com"
	ctx := context.Background()

	expected := &domain.URL{
		Original: URL,
		Short:    "123",
	}

	r.EXPECT().Create(ctx, gomock.Any()).Return(repository.ErrNotUnique)
	r.EXPECT().GetByOriginalURL(ctx, gomock.Any()).Return(expected, nil)

	u, err := s.CreateShortLink(ctx, URL)

	require.Error(t, err)
	assert.Equal(t, URL, u.Original)
}

func TestCreateShortLink_UnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := repository.NewMockRepository(ctrl)
	s := NewService(r)
	ctx := context.Background()
	URL := "https://example.com"

	unexpectedErr := errors.New("unexpected error")

	r.EXPECT().
		Create(ctx, gomock.Any()).
		Return(unexpectedErr)

	result, err := s.CreateShortLink(ctx, URL)

	assert.Error(t, err)
	assert.Equal(t, unexpectedErr, err)
	assert.Nil(t, result)
}

func TestGetOriginalLink_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := repository.NewMockRepository(ctrl)
	s := NewService(r)
	ctx := context.Background()
	shortCode := "nonexistent"

	r.EXPECT().
		Get(ctx, shortCode).
		Return(nil, repository.ErrNotFound)

	result, err := s.GetOriginalLink(ctx, shortCode)

	assert.Error(t, err)
	assert.Equal(t, repository.ErrNotFound, err)
	assert.Empty(t, result)
}

func TestGetOriginalLink_UnexpectedError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := repository.NewMockRepository(ctrl)
	s := NewService(r)
	ctx := context.Background()
	short := "abc123"

	unexpectedErr := errors.New("database error")

	r.EXPECT().
		Get(ctx, short).
		Return(nil, unexpectedErr)

	result, err := s.GetOriginalLink(ctx, short)

	assert.Error(t, err)
	assert.Equal(t, unexpectedErr, err)
	assert.Empty(t, result)
}

func TestGenerateShort_Success(t *testing.T) {
	length := 10

	result, err := generateShort(length)

	require.NoError(t, err)
	assert.Len(t, result, length)

	for _, char := range result {
		assert.Contains(t, alphabet, string(char))
	}
}

func TestGenerateShort_Length(t *testing.T) {
	l := 0

	result, err := generateShort(l)

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Contains(t, err.Error(), "length must be positive")
}
