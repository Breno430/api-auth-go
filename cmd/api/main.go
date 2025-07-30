package main

import (
	"log"

	"api-auth-go/internal/infrastructure/config"
	"api-auth-go/internal/infrastructure/database"
	"api-auth-go/internal/infrastructure/server"
)

func main() {
	cfg := config.Load()

	db, err := database.NewConnection(cfg.GetDatabaseURL())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	srv := server.NewServer(cfg, db)
	if err := srv.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
