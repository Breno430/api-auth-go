package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"api-auth-go/internal/domain/usecases"
	"api-auth-go/internal/infrastructure/config"
	infraRepos "api-auth-go/internal/infrastructure/repositories"
	"api-auth-go/internal/presentation/handlers"
	"api-auth-go/internal/presentation/routes"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	router *gin.Engine
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	userRepo := infraRepos.NewUserRepository(db)

	userUseCase := usecases.NewUserUseCase(userRepo)

	userHandler := handlers.NewUserHandler(userUseCase)

	router := routes.SetupRoutes(userHandler)

	return &Server{
		config: cfg,
		db:     db,
		router: router,
	}
}

func (s *Server) Run() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	log.Printf("Server starting on port %s", s.config.Port)
	return s.router.Run(addr)
}
