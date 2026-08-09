package postgres

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testOriginalURL = "https://example.com/page"
	testShortCode   = "abc123"
)

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	t.Run("creates URL", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		url := domain.NewURL(
			testOriginalURL,
			testShortCode,
		)

		mock.ExpectExec(regexp.QuoteMeta(
			`INSERT INTO urls (short, original, expires_at) VALUES ($1,$2, $3)`,
		)).
			WithArgs(
				url.Short,
				url.Original,
				url.ExpiresAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := r.Create(ctx, url)

		require.NoError(t, err)
	})

	t.Run("returns database error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		url := domain.NewURL(
			testOriginalURL,
			testShortCode,
		)

		wantErr := errors.New("database error")

		mock.ExpectExec(regexp.QuoteMeta(
			`INSERT INTO urls (short, original, expires_at) VALUES ($1,$2, $3)`,
		)).
			WithArgs(
				url.Short,
				url.Original,
				url.ExpiresAt,
			).
			WillReturnError(wantErr)

		err := r.Create(ctx, url)

		require.ErrorIs(t, err, wantErr)
	})
}

func TestRepository_Get(t *testing.T) {
	t.Parallel()

	t.Run("returns existing URL", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		createdAt := time.Now().UTC()
		expiresAt := createdAt.Add(time.Hour)

		rows := sqlmock.NewRows([]string{
			"id",
			"short",
			"original",
			"created_at",
			"expires_at",
		}).AddRow(
			1,
			testShortCode,
			testOriginalURL,
			createdAt,
			expiresAt,
		)

		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT id, short, original, created_at, expires_at FROM urls WHERE short = $1`,
		)).
			WithArgs(testShortCode).
			WillReturnRows(rows)

		got, err := r.Get(ctx, testShortCode)

		require.NoError(t, err)
		require.NotNil(t, got)

		assert.Equal(t, testShortCode, got.Short)
		assert.Equal(t, testOriginalURL, got.Original)
	})

	t.Run("returns ErrNotFound when URL does not exist", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT id, short, original, created_at, expires_at FROM urls WHERE short = $1`,
		)).
			WithArgs("missing").
			WillReturnError(sql.ErrNoRows)

		got, err := r.Get(ctx, "missing")

		require.ErrorIs(t, err, repository.ErrNotFound)
		assert.Nil(t, got)
	})

	t.Run("returns database error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		wantErr := errors.New("database error")

		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT id, short, original, created_at, expires_at FROM urls WHERE short = $1`,
		)).
			WithArgs(testShortCode).
			WillReturnError(wantErr)

		got, err := r.Get(ctx, testShortCode)

		require.ErrorIs(t, err, wantErr)
		assert.Nil(t, got)
	})
}

func TestRepository_GetByOriginalURL(t *testing.T) {
	t.Parallel()

	t.Run("returns existing URL", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		createdAt := time.Now().UTC()
		expiresAt := createdAt.Add(time.Hour)

		rows := sqlmock.NewRows([]string{
			"id",
			"short",
			"original",
			"created_at",
			"expires_at",
		}).AddRow(
			1,
			testShortCode,
			testOriginalURL,
			createdAt,
			expiresAt,
		)

		mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT id, short, original, created_at, expires_at
			FROM urls
			WHERE original = $1`,
		)).
			WithArgs(testOriginalURL).
			WillReturnRows(rows)

		got, err := r.GetByOriginalURL(ctx, testOriginalURL)

		require.NoError(t, err)
		require.NotNil(t, got)

		assert.Equal(t, testShortCode, got.Short)
		assert.Equal(t, testOriginalURL, got.Original)
	})

	t.Run("returns ErrNotFound when URL does not exist", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		missingURL := "https://missing.example.com/page"

		mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT id, short, original, created_at, expires_at
			FROM urls
			WHERE original = $1`,
		)).
			WithArgs(missingURL).
			WillReturnError(sql.ErrNoRows)

		got, err := r.GetByOriginalURL(ctx, missingURL)

		require.ErrorIs(t, err, repository.ErrNotFound)
		assert.Nil(t, got)
	})

	t.Run("returns database error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		wantErr := errors.New("database error")

		mock.ExpectQuery(regexp.QuoteMeta(`
			SELECT id, short, original, created_at, expires_at
			FROM urls
			WHERE original = $1`,
		)).
			WithArgs(testOriginalURL).
			WillReturnError(wantErr)

		got, err := r.GetByOriginalURL(ctx, testOriginalURL)

		require.ErrorIs(t, err, wantErr)
		assert.Nil(t, got)
	})
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()

	t.Run("deletes URL", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		mock.ExpectExec(regexp.QuoteMeta(
			`DELETE FROM urls WHERE short = $1`,
		)).
			WithArgs(testShortCode).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := r.Delete(ctx, testShortCode)

		require.NoError(t, err)
	})

	t.Run("returns database error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		wantErr := errors.New("database error")

		mock.ExpectExec(regexp.QuoteMeta(
			`DELETE FROM urls WHERE short = $1`,
		)).
			WithArgs(testShortCode).
			WillReturnError(wantErr)

		err := r.Delete(ctx, testShortCode)

		require.ErrorIs(t, err, wantErr)
	})

	t.Run("returns nil when URL does not exist", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		r, mock := newMockRepository(t)

		mock.ExpectExec(regexp.QuoteMeta(
			`DELETE FROM urls WHERE short = $1`,
		)).
			WithArgs("missing").
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := r.Delete(ctx, "missing")

		require.Error(t, err, repository.ErrNotFound)
	})
}

func newMockRepository(t *testing.T) (*PostgresRepository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	return New(db), mock
}
