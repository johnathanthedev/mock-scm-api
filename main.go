package main

import (
	"fmt"
	"log"
	"scm-api/db"
	"scm-api/internal/api"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := api.StartServer(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
