package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// ConnectPostgres connects to PostgreSQL and returns a *sql.DB instance
func ConnectPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL")
	return db, nil
}
