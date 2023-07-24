package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load .env file", err)
	}
	dbURI := os.Getenv("DB_URI")
	d, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	d.SetMaxIdleConns(10)
	d.SetMaxOpenConns(25)
	d.SetConnMaxIdleTime(2 * time.Minute)
	d.SetConnMaxLifetime(15 * time.Minute)

	return d
}
