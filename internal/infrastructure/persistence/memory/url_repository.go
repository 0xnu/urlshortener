package memory

import (
	"errors"
	"sync"

	"urlshortener/internal/domain/url"
)

// URLRepositoryMemory is an in-memory implementation of the URLRepository interface
type URLRepositoryMemory struct {
	mu    sync.RWMutex
	store map[string]*url.URL
}

// NewURLRepositoryMemory creates a new URLRepositoryMemory
func NewURLRepositoryMemory() *URLRepositoryMemory {
	return &URLRepositoryMemory{
		store: make(map[string]*url.URL),
	}
}

// Save saves a URL into the repository
func (r *URLRepositoryMemory) Save(u *url.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if URL already exists
	if _, exists := r.store[u.ShortCode]; exists {
		return errors.New("ShortCode already exists")
	}

	r.store[u.ShortCode] = u
	return nil
}

// FindByShortCode retrieves a URL by its shortcode
func (r *URLRepositoryMemory) FindByShortCode(shortCode string) (*url.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if u, exists := r.store[shortCode]; exists {
		return u, nil
	}
	return nil, errors.New("ShortCode not found")
}

// DeleteByShortCode deletes a URL by its shortcode
func (r *URLRepositoryMemory) DeleteByShortCode(shortCode string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.store[shortCode]; !exists {
		return errors.New("ShortCode not found")
	}

	delete(r.store, shortCode)
	return nil
}
