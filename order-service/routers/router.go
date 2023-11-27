package routers

import (
	"order-service/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupOrderRoutes sets up the order-related routes
func SetupOrderRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1/orders")
	{
		v1.GET("", handlers.ListOrders)
		v1.GET("/user/:user_id", handlers.ListOrdersByUser)
		v1.GET("/:id", handlers.GetOrder)
		v1.POST("", handlers.CreateOrder)
		v1.DELETE("/limit/:id", handlers.DeleteLimitOrder)
	}

	// Other order-related routes can be added here
}
