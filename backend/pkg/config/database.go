package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB thiết lập kết nối đến PostgreSQL database
func ConnectDB(config *Config) *sql.DB {
	connectionStr := config.GetDatabaseURL()

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test kết nối
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Connected to PostgreSQL database!")
	return db
}
