package main

import (
	"database/sql"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"

	serverPkg "gocrud/internal/http"
	"gocrud/internal/stor"
	"gocrud/internal/utils/logger"
)

// @title           Go CRUD Storage API
// @version         1.0
// @description     Simple CRUD API on Go + chi + Postgres

// @contact.name   API Support
// @contact.email  support@example.com

// @BasePath  /

func init() {
	logger.SetLogger(os.Getenv("LOGGER_ENV"))
}

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

	log.Info("Successfully open connect with DB")

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("cannot connect to postgres: %v", err)
	}

	log.Println("Connected to PostgreSQL")

	storage := stor.NewStorage(db)
	httpServer := serverPkg.CreatServer(storage)

	log.Info(fmt.Sprintf("Server is startingдl isten on port %d", 8080))

	if err := httpServer.Listen(); err != nil {
		log.Fatalf("server stopped with error: %v", err)
	}
}
