package main

import (
	"fmt"
	"log"
	"net/http"

	app "urlshortener/internal/app/urlshortener"
	cfg "urlshortener/internal/config"
	httpAPI "urlshortener/internal/infrastructure/api/http"
	sqlRepo "urlshortener/internal/infrastructure/persistence/sql"
)

func main() {
	// Load configuration
	config := cfg.LoadConfig()

	// Initialize the MySQL URL repository
	dsn := config.GetMySQLDSN()
	urlRepo, err := sqlRepo.NewURLRepositorySQL(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize URL repository: %v", err)
	}

	// Initialize the core application logic
	application := app.NewApplication(urlRepo)

	// Initialize and start the HTTP server
	router := httpAPI.NewRouter(application)
	http.Handle("/", router.SetupRoutes())
	log.Printf("Server started at :%d", config.HTTPPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.HTTPPort), nil))

}
