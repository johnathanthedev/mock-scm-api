package main

import (
	"fmt"
	"log"
	"scm-api/api"
	"scm-api/db"
	ws "scm-api/ws"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
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
