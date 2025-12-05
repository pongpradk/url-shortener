package service

import (
	"context"
	"crypto/md5"
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

// generateUniqueID creates a unique ID from timestamp and input string
func generateUniqueID(input string) uint64 {
	uniqueID := uint64(time.Now().UnixNano())

	// Mix in hash of input for extra randomness
	hash := md5.Sum([]byte(input + time.Now().String()))
	for i := range 8 {
		uniqueID ^= uint64(hash[i]) << (i * 8)
	}

	return uniqueID
}

// ShortenURL creates a short URL from long URL
func (s *URLService) ShortenURL(ctx context.Context, longURL string) (string, error) {
	// Check if URL already exists
	existing, err := s.repo.FindByLongURL(ctx, longURL)
	if err == nil {
		return existing.ShortURL, nil
	}

	if !errors.Is(err, repository.ErrNotFound) {
		return "", err
	}

	uniqueID := generateUniqueID(longURL)

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
