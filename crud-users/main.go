// ./crud-users/main.go

package main

import (
	"crud-users/docs"
	"crud-users/models"
	"crud-users/routers"
	"fmt"
	"log"
	"os"

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

	// Setup user routes
	routers.SetupUserRoutes(r, db)

	// Setup Swagger documentation route
	//routers.SetupSwaggerRoute(r)

	return r
}

func main() {
	// Get containerName from environment variable
	containerName := os.Getenv("CONTAINER_DB")

	// Check if containerName is empty
	if containerName == "" {
		log.Println("CONTAINER_DB environment variable is not set.")
		return
	}

	// Format the dsn string with the containerName
	dsn := fmt.Sprintf("root:root@tcp(%s:3306)/db_americas_technology?charset=utf8mb4&parseTime=True&loc=Local", containerName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to the database.")
		return
	}

	// AutoMigrate will create the necessary table. You can also use CreateTable to only create the table if it does not exist.
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Failed to migrate the database.")
		return
	}

	swaggerURL := "http://localhost:8080/swagger/doc.json"

	r := SetupRouter(db)

	url := ginSwagger.URL(swaggerURL)

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Run the application
	r.Run(":8080")
}
