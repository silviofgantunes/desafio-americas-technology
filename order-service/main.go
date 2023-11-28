package main

import (
	"fmt"
	"order-service/docs"
	"order-service/models"
	"order-service/routers"

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

	// Setup order routes
	routers.SetupOrderRoutes(r, db)

	// Setup Swagger documentation route
	//url := ginSwagger.URL("http://localhost:8081/swagger/doc.json")
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}

func main() {
	dsn := "root:root@tcp(database-americas-technology:3306)/db_americas_technology?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate will create the necessary table. You can also use CreateTable to only create the table if it does not exist.
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		panic("Failed to migrate the database")
	}

	hostIP := "host.docker.internal"

	swaggerURL := fmt.Sprintf("http://%s:8081/swagger/doc.json", hostIP)

	r := SetupRouter(db)

	url := ginSwagger.URL(swaggerURL)

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Run the application
	r.Run(":8081")
}
