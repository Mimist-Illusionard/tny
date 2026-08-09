package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
	"github.com/jackc/pgx/v5/pgconn"
)

const uniqueViolationCode = "23505"

type PostgresRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, url *domain.URL) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO urls (short, original, expires_at) VALUES ($1, $2, $3)`,
		url.Short, url.Original, url.ExpiresAt)
	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationCode {
		switch pgErr.ConstraintName {
		case "urls_short_unique":
			return repository.ErrExists
		case "urls_original_unique":
			return repository.ErrNotUnique
		}
	}

	return fmt.Errorf("create url: %w", err)
}

func (r *PostgresRepository) Get(ctx context.Context, short string) (*domain.URL, error) {
	u := &domain.URL{}

	err := r.db.QueryRowContext(ctx,
		`SELECT id, short, original, created_at, expires_at FROM urls WHERE short = $1`, short).
		Scan(&u.ID, &u.Short, &u.Original, &u.CreatedAt, &u.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("get url by short code: %w", err)
	}

	return u, nil
}

func (r *PostgresRepository) GetByOriginalURL(ctx context.Context, originalURL string) (*domain.URL, error) {
	u := &domain.URL{}

	err := r.db.QueryRowContext(ctx, `
		SELECT id, short, original, created_at, expires_at
		FROM urls
		WHERE original = $1`, originalURL).
		Scan(&u.ID, &u.Short, &u.Original, &u.CreatedAt, &u.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("get url by original url: %w", err)
	}

	return u, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, short string) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM urls WHERE short = $1`, short)
	if err != nil {
		return fmt.Errorf("delete url: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get deleted rows count: %w", err)
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
