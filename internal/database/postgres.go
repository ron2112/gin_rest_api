package database

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	ctx := context.Background() // context.Context
	config, err := pgxpool.ParseConfig(databaseURL) // pgxpool.Pool (v5)
	if err != nil {
		log.Printf("Unable to parse DATABASE_URL: %v", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Unable to create connection pool: %v", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("Unable to ping database: %v", err)
		pool.Close()
		return nil, err
	}
	log.Println("Successfully connected to Postgresql database")
	return pool, nil
}