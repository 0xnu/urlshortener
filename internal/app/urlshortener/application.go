package urlshortener

import (
	"math/rand"
	"time"

	"urlshortener/internal/domain/url"
)

// Application holds the domain services and repositories
type Application struct {
	urlRepo url.URLRepository
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// NewApplication creates a new Application
func NewApplication(urlRepo url.URLRepository) *Application {
	return &Application{
		urlRepo: urlRepo,
	}
}

func generateShortCode() string {
	rand.Seed(time.Now().UnixNano())
	length := 5
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
