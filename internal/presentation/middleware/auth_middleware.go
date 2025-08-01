package middleware

import (
	"net/http"
	"strings"

	"api-auth-go/internal/infrastructure/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format. Use 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_name", claims.Name)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("user_role")
		if userRole != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied. Admin role required",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RoleBasedAccessMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("user_role")
		userID := c.GetString("user_id")
		requestedUserID := c.Param("id")

		if requestedUserID != "" {
			if userRole == "admin" {
				c.Next()
				return
			}

			if userID != requestedUserID {
				c.JSON(http.StatusForbidden, gin.H{
					"error": "Access denied. You can only access your own data",
				})
				c.Abort()
				return
			}

			if c.Request.Method == "DELETE" {
				if userRole == "user" {
					c.JSON(http.StatusForbidden, gin.H{
						"error": "Access denied. Users cannot delete themselves",
					})
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}
