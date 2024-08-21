package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"urlshortener/internal/app/urlshortener"
)

// Handlers holds the application logic and router
type Handlers struct {
	app *urlshortener.Application
}

// NewHandlers creates a new Handlers instance
func NewHandlers(app *urlshortener.Application) *Handlers {
	return &Handlers{
		app: app,
	}
}

// ShortenURLHandler handles URL shortening
func (h *Handlers) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		OriginalURL string `json:"original_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortCode, err := h.app.ShortenURL(payload.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fullShortURL := fmt.Sprintf("http://localhost/%s", shortCode)

	response := struct {
		ShortCode    string `json:"short_code"`
		FullShortURL string `json:"full_short_url"`
	}{
		ShortCode:    shortCode,
		FullShortURL: fullShortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RedirectURLHandler handles URL retrieval and redirection
func (h *Handlers) RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[len("/"):]

	originalURL, err := h.app.GetOriginalURL(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}

// DeleteURLHandler handles URL deletion
func (h *Handlers) DeleteURLHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[len("/delete/"):]

	if err := h.app.DeleteURL(shortCode); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
