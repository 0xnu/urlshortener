package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config holds the configuration values for the application
type Config struct {
	HTTPPort      int
	MySQLUser     string
	MySQLPassword string
	MySQLHost     string
	MySQLPort     int
	MySQLDatabase string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *Config {
	var cfg Config

	// Load HTTP port
	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		log.Fatal("HTTP_PORT environment variable is required")
	}
	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		log.Fatalf("Invalid HTTP_PORT: %v", err)
	}
	cfg.HTTPPort = httpPort

	// Load MySQL User
	cfg.MySQLUser = os.Getenv("MYSQL_USER")
	if cfg.MySQLUser == "" {
		log.Fatal("MYSQL_USER environment variable is required")
	}

	// Load MySQL Password
	cfg.MySQLPassword = os.Getenv("MYSQL_PASSWORD")
	if cfg.MySQLPassword == "" {
		log.Println("MYSQL_PASSWORD environment variable is not set, using default")
	}

	// Load MySQL Host
	cfg.MySQLHost = os.Getenv("MYSQL_HOST")
	if cfg.MySQLHost == "" {
		log.Fatal("MYSQL_HOST environment variable is required")
	}

	// Load MySQL Port
	mysqlPortStr := os.Getenv("MYSQL_PORT")
	if mysqlPortStr == "" {
		log.Fatal("MYSQL_PORT environment variable is required")
	}
	mysqlPort, err := strconv.Atoi(mysqlPortStr)
	if err != nil {
		log.Fatalf("Invalid MYSQL_PORT: %v", err)
	}
	cfg.MySQLPort = mysqlPort

	// Load MySQL Database
	cfg.MySQLDatabase = os.Getenv("MYSQL_DATABASE")
	if cfg.MySQLDatabase == "" {
		log.Fatal("MYSQL_DATABASE environment variable is required")
	}

	return &cfg
}

// GetMySQLDSN generates the MySQL DSN based on the loaded configuration
func (cfg *Config) GetMySQLDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.MySQLUser,
		cfg.MySQLPassword,
		cfg.MySQLHost,
		cfg.MySQLPort,
		cfg.MySQLDatabase,
	)
}
