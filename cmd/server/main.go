package main

import (
	"log"
	"net/http"

	"github.com/homecloud/go-api/internal/config"
	"github.com/homecloud/go-api/internal/database"
	"github.com/homecloud/go-api/internal/health"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", health.HealthHandler)

	addr := ":8080"
	log.Printf("starting server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
