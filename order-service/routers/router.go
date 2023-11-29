// ./order-service/routers/router.go

package routers

import (
	"net/http"
	"order-service/handlers"
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

// SetupOrderRoutes sets up the order-related routes
func SetupOrderRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1/orders")
	v1.Use(AuthMiddleware())
	{
		v1.GET("", handlers.ListOrders)
		v1.GET("/user/:user_id", handlers.ListOrdersByUser)
		v1.GET("/:id", handlers.GetOrder)
		v1.POST("", handlers.CreateOrder)
		v1.DELETE("/limit/:id", handlers.DeleteLimitOrder)
	}

	// Other order-related routes can be added here
}
