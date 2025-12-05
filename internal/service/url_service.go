package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"github.com/pongpradk/url-shortener/internal/encoder"
	"github.com/pongpradk/url-shortener/internal/repository"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

// ShortenURL creates a short URL from long URL
func (s *URLService) ShortenURL(ctx context.Context, longURL string) (string, error) {
	// Check if URL already exists
	existing, err := s.repo.FindByLongURL(ctx, longURL)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return "", err
	}

	if existing != nil {
		return existing.ShortURL, nil
	}

	// Generate unique ID from URL + timestamp
	hash := md5.Sum([]byte(longURL + time.Now().String()))
	hashStr := hex.EncodeToString(hash[:])

	// Convert first 8 bytes to uint64
	var uniqueID uint64
	for i := 0; i < 8 && i < len(hashStr); i++ {
		uniqueID = uniqueID*256 + uint64(hashStr[i])
	}

	// Generate short URL
	shortURL := encoder.Encode(uniqueID)

	// Save to database
	_, err = s.repo.Create(ctx, shortURL, longURL)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

// GetLongURL retrieves original URL from short URL
func (s *URLService) GetLongURL(ctx context.Context, shortURL string) (string, error) {
	url, err := s.repo.FindByShortURL(ctx, shortURL)
	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}
