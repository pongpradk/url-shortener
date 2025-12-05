package repository

import (
	"database/sql"
	"errors"
)

type URL struct {
	ID       int64
	ShortURL string
	LongURL  string
}

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

// FindByLongURL checks if long URL already exists
func (r *URLRepository) FindByLongURL(longURL string) (*URL, error) {
	var url URL
	query := "SELECT id, short_url, long_url FROM urls WHERE long_url = $1"

	err := r.db.QueryRow(query, longURL).Scan(&url.ID, &url.ShortURL, &url.LongURL)
	if err == sql.ErrNoRows {
		return nil, nil // Not found
	}
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// FindByShortURL retrieves long URL by short URL
func (r *URLRepository) FindByShortURL(shortURL string) (*URL, error) {
	var url URL
	query := "SELECT id, short_url, long_url FROM urls WHERE short_url = $1"

	err := r.db.QueryRow(query, shortURL).Scan(&url.ID, &url.ShortURL, &url.LongURL)
	if err == sql.ErrNoRows {
		return nil, errors.New("URL not found")
	}
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// Create inserts a new URL mapping
func (r *URLRepository) Create(shortURL, longURL string) (*URL, error) {
	var id int64
	query := "INSERT INTO urls (short_url, long_url) VALUES ($1, $2) RETURNING id"

	err := r.db.QueryRow(query, shortURL, longURL).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &URL{
		ID:       id,
		ShortURL: shortURL,
		LongURL:  longURL,
	}, nil
}
