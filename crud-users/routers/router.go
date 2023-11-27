package routers

// @title CRUD Users
// @version 1.0
// @description Serviço básico de CRUD de usuários.

import (
	"crud-users/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupUserRoutes sets up the user-related routes
func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/api/v1/users")
	{
		v1.GET("/:id", handlers.GetUser)
		v1.POST("", handlers.CreateUser)
		v1.PUT("/:id", handlers.UpdateUser)
		v1.DELETE("/:id", handlers.DeleteUser)
		v1.GET("", handlers.ListUsers)
	}

	// Other user-related routes can be added here
}
