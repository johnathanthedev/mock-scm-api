package main

import (
	"log"
	"os"
	"scm-api/api"
	"scm-api/db"
	ws "scm-api/ws"

	"github.com/joho/godotenv"
)

func main() {
	// Attempt to load the .env file if it exists, but don't exit if it's missing.
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Println("Error loading .env file", err)
	}

	broker := ws.NewBroker()
	go broker.Run()

	if err := db.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := api.StartServer(broker); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
