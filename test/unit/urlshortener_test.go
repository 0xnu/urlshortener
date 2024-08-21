package urlshortener_test

import (
	"testing"

	"urlshortener/internal/app/urlshortener"
	"urlshortener/internal/infrastructure/persistence/memory"
)

func TestShortenURL(t *testing.T) {
	repo := memory.NewURLRepositoryMemory()
	app := urlshortener.NewApplication(repo)

	tests := []struct {
		name        string
		originalURL string
		wantErr     bool
	}{
		{"Valid URL", "https://google.com", false},
		{"Empty URL", "", true},
		{"Invalid URL", "invalid-url", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := app.ShortenURL(tt.originalURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortenURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOriginalURL(t *testing.T) {
	repo := memory.NewURLRepositoryMemory()
	app := urlshortener.NewApplication(repo)

	// Setup: Shorten a URL for testing
	shortCode, _ := app.ShortenURL("https://google.com")

	tests := []struct {
		name      string
		shortCode string
		wantErr   bool
	}{
		{"Valid shortcode", shortCode, false},
		{"Empty shortcode", "", true},
		{"Non-existent shortcode", "nonexistent", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := app.GetOriginalURL(tt.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOriginalURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteURL(t *testing.T) {
	repo := memory.NewURLRepositoryMemory()
	app := urlshortener.NewApplication(repo)

	// Setup: Shorten a URL for testing
	shortCode, _ := app.ShortenURL("https://google.com")

	tests := []struct {
		name      string
		shortCode string
		wantErr   bool
	}{
		{"Valid shortcode", shortCode, false},
		{"Empty shortcode", "", true},
		{"Non-existent shortcode", "nonexistent", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := app.DeleteURL(tt.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
