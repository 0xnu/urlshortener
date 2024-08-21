package url

import (
	"errors"
	"time"
)

// URL represents a shortened URL
type URL struct {
	ShortCode   string
	OriginalURL string
	CreatedAt   time.Time
	ExpiresAt   time.Time
}

// NewURL creates a new URL instance and validates it
func NewURL(shortCode, originalURL string, expiresAt time.Time) (*URL, error) {
	if shortCode == "" || originalURL == "" {
		return nil, errors.New("shortCode and originalURL must not be empty")
	}

	if expiresAt.Before(time.Now()) {
		return nil, errors.New("expiresAt must be in the future")
	}

	return &URL{
		ShortCode:   shortCode,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
		ExpiresAt:   expiresAt,
	}, nil
}

// IsExpired checks if the URL is expired
func (u *URL) IsExpired() bool {
	return time.Now().After(u.ExpiresAt)
}

// URLRepository defines the interface for a URL repository
type URLRepository interface {
	// Save saves a URL into the repository
	Save(u *URL) error

	// FindByShortCode retrieves a URL by its shortcode
	FindByShortCode(shortCode string) (*URL, error)

	// DeleteByShortCode deletes a URL by its shortcode
	DeleteByShortCode(shortCode string) error
}
