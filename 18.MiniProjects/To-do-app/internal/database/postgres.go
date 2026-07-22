package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Printf("Failed to parse database URL: %v", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Failed to create connection pool: %v", err)
		return nil, err
	}

	return pool, nil
}
