package main

import (
	"context"
	"log"

	"api-auth-go/internal/domain/entities"
	"api-auth-go/internal/infrastructure/config"
	"api-auth-go/internal/infrastructure/database"
	"api-auth-go/internal/infrastructure/repositories"
	"api-auth-go/internal/infrastructure/server"

	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := database.NewConnection(cfg.GetDatabaseURL())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := runSeed(db); err != nil {
		log.Printf("Warning: Failed to run seed: %v", err)
	}

	srv := server.NewServer(cfg, db)
	if err := srv.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func runSeed(db *gorm.DB) error {
	userRepo := repositories.NewUserRepository(db)

	ctx := context.Background()
	adminEmail := "admin@example.com"

	existingAdmin, err := userRepo.FindByEmail(ctx, adminEmail)
	if err != nil {
		return err
	}

	if existingAdmin != nil {
		log.Println("Admin user already exists, skipping seed")
		return nil
	}

	adminPassword := "admin123"
	admin, err := entities.NewAdminUser("Admin User", adminEmail, adminPassword)
	if err != nil {
		return err
	}

	if err := userRepo.Create(ctx, admin); err != nil {
		return err
	}

	log.Println("✅ Admin user created successfully!")
	log.Printf("📧 Email: %s", adminEmail)
	log.Printf("🔑 Password: %s", adminPassword)
	log.Printf("👑 Role: admin")
	log.Println("💡 Use this admin to create other users via API")

	return nil
}
