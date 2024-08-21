package http

import (
	"net/http"
	"urlshortener/internal/app/urlshortener"
)

// Router holds the handlers and provides methods for route setup
type Router struct {
	handlers *Handlers
}

// NewRouter creates a new Router instance
func NewRouter(app *urlshortener.Application) *Router {
	return &Router{
		handlers: NewHandlers(app),
	}
}

// SetupRoutes sets up the routes for the HTTP server
func (r *Router) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Route for shortening URLs
	mux.HandleFunc("/shorten", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodPost {
			r.handlers.ShortenURLHandler(w, req)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route for redirecting to original URLs
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet && req.URL.Path != "/" {
			r.handlers.RedirectURLHandler(w, req)
		} else {
			http.Error(w, "Method not allowed or invalid path", http.StatusMethodNotAllowed)
		}
	})

	// Route for deleting URLs
	mux.HandleFunc("/delete/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodDelete {
			r.handlers.DeleteURLHandler(w, req)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
