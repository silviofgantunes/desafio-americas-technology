// ./auth-service/routers/router.go

package routers

import (
	"auth-service/handlers"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes configura as rotas de autenticação
func SetupAuthRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1/auth")
	{
		// Rotas existentes
		v1.GET("/generate-token", handlers.GenerateToken)
		v1.POST("/admins", handlers.CreateAdmin)
		v1.GET("/admins", handlers.ListAdmins)
		v1.GET("/admins/:id", handlers.GetAdmin)
		v1.PUT("/admins/:id", handlers.UpdateAdmin)
		v1.DELETE("/admins/:id", handlers.DeleteAdmin)
	}
}
