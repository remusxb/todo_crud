package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	// Database connection string
	connStr := "postgres://postgres:postgres@db/main?sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	slog.Info("Successfully connected to the database")
	return db, nil
}
