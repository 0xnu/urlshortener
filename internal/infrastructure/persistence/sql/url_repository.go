package sql

import (
	"database/sql"
	"errors"

	"urlshortener/internal/domain/url"

	// Import MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// URLRepositorySQL is a MySQL implementation of the URLRepository interface
type URLRepositorySQL struct {
	db *sql.DB
}

// NewURLRepositorySQL creates a new URLRepositorySQL
func NewURLRepositorySQL(dsn string) (*URLRepositorySQL, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &URLRepositorySQL{db: db}, nil
}

// Save saves a URL into the repository
func (r *URLRepositorySQL) Save(u *url.URL) error {
	query := "INSERT INTO urls (short_code, original_url) VALUES (?, ?)"
	_, err := r.db.Exec(query, u.ShortCode, u.OriginalURL)
	if err != nil {
		return errors.New("ShortCode already exists")
	}
	return nil
}

// FindByShortCode retrieves a URL by its shortcode
func (r *URLRepositorySQL) FindByShortCode(shortCode string) (*url.URL, error) {
	query := "SELECT short_code, original_url FROM urls WHERE short_code = ?"
	row := r.db.QueryRow(query, shortCode)

	var u url.URL
	err := row.Scan(&u.ShortCode, &u.OriginalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ShortCode not found")
		}
		return nil, err
	}
	return &u, nil
}

// DeleteByShortCode deletes a URL by its shortcode
func (r *URLRepositorySQL) DeleteByShortCode(shortCode string) error {
	query := "DELETE FROM urls WHERE short_code = ?"
	_, err := r.db.Exec(query, shortCode)
	if err != nil {
		return errors.New("ShortCode not found")
	}
	return nil
}
