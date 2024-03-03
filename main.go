package main

import (
	"log"
	"scm-api/internal/api"
)

func main() {
    if err := api.StartServer(); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}
