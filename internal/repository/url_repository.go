package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("record not found")
)

type URL struct {
	ID        int64
	ShortURL  string
	LongURL   string
	CreatedAt time.Time
}

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

// FindByLongURL checks if long URL already exists
func (r *URLRepository) FindByLongURL(ctx context.Context, longURL string) (*URL, error) {
	var url URL
	query := "SELECT id, short_url, long_url, created_at FROM urls WHERE long_url = $1"

	err := r.db.QueryRowContext(ctx, query, longURL).Scan(
		&url.ID,
		&url.ShortURL,
		&url.LongURL,
		&url.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &url, nil
}

// FindByShortURL retrieves long URL by short URL
func (r *URLRepository) FindByShortURL(ctx context.Context, shortURL string) (*URL, error) {
	var url URL
	query := "SELECT id, short_url, long_url, created_at FROM urls WHERE short_url = $1"

	err := r.db.QueryRowContext(ctx, query, shortURL).Scan(
		&url.ID,
		&url.ShortURL,
		&url.LongURL,
		&url.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &url, nil
}

// Create inserts a new URL mapping
func (r *URLRepository) Create(ctx context.Context, shortURL, longURL string) (*URL, error) {
	var url URL
	query := `
		INSERT INTO urls (short_url, long_url) 
		VALUES ($1, $2) 
		RETURNING id, short_url, long_url, created_at
	`

	err := r.db.QueryRowContext(ctx, query, shortURL, longURL).Scan(
		&url.ID,
		&url.ShortURL,
		&url.LongURL,
		&url.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
