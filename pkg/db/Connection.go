package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func GetConnection() *pgxpool.Pool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load .env file")
	}
	dbURI := os.Getenv("DB_URI")
	p, err2 := pgxpool.New(context.Background(), dbURI)
	if err2 != nil {
		panic(err2)
	}

	return p
}
