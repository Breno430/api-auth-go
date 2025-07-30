package routes

import (
	"github.com/gin-gonic/gin"

	"api-auth-go/internal/presentation/handlers"
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
	}

	return router
}
