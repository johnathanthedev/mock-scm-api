package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//lint:ignore U1000 main is intentionally left unused as it serves as the entry point for this file
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	m, err := migrate.New(
		"file://db/migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatalf("could not create migration instance: %v", err)
	}

	const rollbackSteps = 1

	if err := m.Steps(-rollbackSteps); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("no migrations to rollback")
		} else {
			log.Fatalf("could not rollback migrations: %v", err)
		}
	}

	log.Printf("%d migration(s) rolled back successfully", rollbackSteps)
}
