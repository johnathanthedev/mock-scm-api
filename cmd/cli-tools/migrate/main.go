package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

//lint:ignore U1000 main is intentionally left unused as it serves as the entry point for this file
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	databaseURL := os.Getenv("POSTGRESQL_URL")
	if databaseURL == "" {
		log.Fatal("POSTGRESQL_URL environment variable is not set")
	}

	m, err := migrate.New(
		"file://db/migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatalf("could not create migration instance: %v", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("no migrations to apply")
		} else {
			log.Fatalf("could not apply migrations: %v", err)
		}
	}

	log.Println("Migrations applied successfully")
}
