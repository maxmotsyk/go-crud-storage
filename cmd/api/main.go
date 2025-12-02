package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	serverPkg "gocrud/internal/http"
	"gocrud/internal/stor"
)

func main() {
	// Берём DSN только из окружения
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("cannot open db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("cannot connect to postgres: %v", err)
	}

	log.Println("Connected to PostgreSQL")

	storage := stor.NewStorage(db)
	httpServer := serverPkg.CreatServer(storage)

	log.Println("Server listening on :8080")

	if err := httpServer.Listen(); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}
