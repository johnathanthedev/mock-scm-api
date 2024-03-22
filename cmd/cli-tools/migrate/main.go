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
