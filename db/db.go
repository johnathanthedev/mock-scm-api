package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Connect() error {
	psqlURL := os.Getenv("POSTGRESQL_URL")
	if psqlURL == "" {
		return fmt.Errorf("POSTGRESQL_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(psqlURL), &gorm.Config{})
	if err != nil {
		return err
	}

	database = db
	fmt.Println("Database connection established")
	return nil
}

func Close() {
	if database != nil {
		db, _ := database.DB()
		db.Close()
	}
}
