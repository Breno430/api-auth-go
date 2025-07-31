package routes

import (
	"github.com/gin-gonic/gin"

	"api-auth-go/internal/infrastructure/services"
	"api-auth-go/internal/presentation/handlers"
	"api-auth-go/internal/presentation/middleware"
)

func SetupRoutes(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	healthHandler := handlers.NewHealthHandler()
	router.GET("/health", healthHandler.HealthCheck)

	userRoutes := router.Group("/api/v1/users")
	{
		userRoutes.POST("/signup", userHandler.CreateUser)
		userRoutes.POST("/login", userHandler.Login)
	}

	passwordResetRoutes := router.Group("/api/v1/password-reset")
	{
		passwordResetRoutes.POST("/request", userHandler.RequestPasswordReset)
		passwordResetRoutes.POST("/reset", userHandler.ResetPassword)
	}

	jwtService := services.NewJWTService()
	protectedRoutes := router.Group("/api/v1")
	protectedRoutes.Use(middleware.AuthMiddleware(jwtService))
	{
		protectedRoutes.GET("/profile", userHandler.GetProfile)
	}

	return router
}
