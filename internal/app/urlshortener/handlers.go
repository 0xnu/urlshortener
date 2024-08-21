package urlshortener

import (
	"time"

	"urlshortener/internal/domain/url"
)

// ShortenURL shortens a given URL
func (app *Application) ShortenURL(originalURL string) (string, error) {
	if originalURL == "" {
		return "", NewAppError(ErrorTypeValidation, "Original URL cannot be empty")
	}

	// Generate a shortcode (for simplicity, using a timestamp-based approach)
	shortCode := generateShortCode()

	// Create a new URL with an expiration time of 24 hours from now
	expiresAt := time.Now().Add(24 * time.Hour)
	newURL, err := url.NewURL(shortCode, originalURL, expiresAt)
	if err != nil {
		return "", NewAppError(ErrorTypeValidation, err.Error())
	}

	// Save the URL to the repository
	if err := app.urlRepo.Save(newURL); err != nil {
		return "", NewAppError(ErrorTypeDatabase, "Failed to save URL")
	}

	return shortCode, nil
}

// GetOriginalURL retrieves the original URL by its shortcode
func (app *Application) GetOriginalURL(shortCode string) (string, error) {
	u, err := app.urlRepo.FindByShortCode(shortCode)
	if err != nil {
		return "", NewAppError(ErrorTypeNotFound, "ShortCode not found")
	}

	if u.IsExpired() {
		return "", NewAppError(ErrorTypeNotFound, "URL has expired")
	}

	return u.OriginalURL, nil
}

// DeleteURL deletes a URL by its shortcode
func (app *Application) DeleteURL(shortCode string) error {
	if err := app.urlRepo.DeleteByShortCode(shortCode); err != nil {
		return NewAppError(ErrorTypeNotFound, "ShortCode not found")
	}
	return nil
}
