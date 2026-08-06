package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Mimist-Illusionard/url-shortener/internal/domain"
	"github.com/Mimist-Illusionard/url-shortener/internal/repository"
)

type PostgresRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, url *domain.URL) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO urls (short, original, expiresAt) VALUES (?, ?, ?)`,
		url.Short, url.Original, url.ExpiresAt)

	return err
}

func (r *PostgresRepository) Get(ctx context.Context, short string) (*domain.URL, error) {
	url := &domain.URL{}

	row := r.db.QueryRowContext(ctx,
		`SELECT id, short, original, createdAt, expiresAt FROM urls WHERE short = ?`, short)

	err := row.Scan(&url.ID, &url.Short, &url.Original, &url.CreatedAt, &url.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return url, nil
}

func (r *PostgresRepository) GetByOriginalURL(ctx context.Context, url string) (*domain.URL, error) {
	u := &domain.URL{}

	row := r.db.QueryRowContext(ctx, `
		SELECT id, short, original, createdAt, expiresAt 
		FROM urls 
		WHERE original = ?`, url)

	err := row.Scan(&u.ID, &u.Short, &u.Original, &u.CreatedAt, &u.ExpiresAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, short string) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM urls WHERE short = ?`, short)

	return err
}
