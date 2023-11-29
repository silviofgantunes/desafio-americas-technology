// ./crud-users/routers/router.go

package routers

// @title CRUD Users
// @version 1.0
// @description Serviço básico de CRUD de usuários.

import (
	"crud-users/handlers"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]

		// Obter a chave secreta da variável de ambiente
		secretKey := os.Getenv("JWT_SECRET_KEY")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// SetupUserRoutes sets up the user-related routes
func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1/users")
	v1.Use(AuthMiddleware())
	{
		v1.GET("/:id", handlers.GetUser)
		v1.POST("", handlers.CreateUser)
		v1.PUT("/:id", handlers.UpdateUser)
		v1.DELETE("/:id", handlers.DeleteUser)
		v1.GET("", handlers.ListUsers)
		//v1.GET("/check-user/:id", handlers.CheckUser)
	}

	// Other user-related routes can be added here
}
