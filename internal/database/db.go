package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect(){
	url := "postgresql://postgres:postgres@localhost:5432/postgres"

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %vn", err)
	}
	DB = pool
	log.Println("Connected to PostgreSQL")
}
