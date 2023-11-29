package main

import (
	"auth-service/docs"
	"auth-service/models"
	"auth-service/routers"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Setup authentication routes
	routers.SetupAuthRoutes(r)

	// Run the authentication service
	r.Run(":8082")

	return r
}

func main() {
	dsn := "root:root@tcp(database-americas-technology:3306)/db_americas_technology?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate will create the necessary table. You can also use CreateTable to only create the table if it does not exist.
	err = db.AutoMigrate(&models.Admin{})
	if err != nil {
		panic("Failed to migrate the database")
	}

	hostIP := "host.docker.internal"

	swaggerURL := fmt.Sprintf("http://%s:8082/swagger/doc.json", hostIP)

	r := SetupRouter(db)

	url := ginSwagger.URL(swaggerURL)

	docs.SwaggerInfo.BasePath = "/api/v1/auth"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Run the application
	r.Run(":8082")
}
