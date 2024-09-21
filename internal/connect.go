package internal

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

// ConnectDB establishes a connection to the PostgreSQL database using pgx
func ConnectDB(dbURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse DB URL: %v", err)
		return nil, err
	}

	// Setting maximum connection pool size (optional)
	config.MaxConns = 10

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database!")
	return pool, nil
}

// CloseDB is used to close the connection to the database
func CloseDB(conn *pgxpool.Pool) {
	conn.Close()
	log.Println("Database connection closed.")
}
