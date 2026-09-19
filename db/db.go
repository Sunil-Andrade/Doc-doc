package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "postgres://postgres:postgres@localhost:5432/docdoc"
	}

	conn, err := pgx.Connect(
		context.Background(),
		databaseURL,
	)

	if err != nil {
		log.Fatal("database connection failed:", err)
	}

	log.Println("database connected")

	return conn
}
