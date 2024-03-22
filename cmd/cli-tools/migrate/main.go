package main

import (
	"flag"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	var forceVersion int
	flag.IntVar(&forceVersion, "force", 0, "Force the version to this value if there are issues")
	flag.Parse()

	// Attempt to load the .env file if it exists, but don't exit if it's missing.
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Println("Error loading .env file", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if migrationsPath == "" {
		// Fallback to a local path if MIGRATIONS_PATH is not set
		migrationsPath = "file://db/migrations"
	}

	log.Printf("Using migrations path: %s", migrationsPath)
	m, err := migrate.New(migrationsPath, databaseURL)
	if err != nil {
		log.Fatalf("could not create migration instance: %v", err)
	}

	if forceVersion != 0 {
		if err := m.Force(forceVersion); err != nil {
			log.Fatalf("could not force version %d: %v", forceVersion, err)
		}
		log.Printf("Forced version to %d successfully", forceVersion)
		return
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("no migrations to apply")
		} else {
			log.Fatalf("could not apply migrations: %v", err)
		}
	} else {
		log.Println("Migrations applied successfully")
	}
}
