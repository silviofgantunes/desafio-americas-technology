package main

import (
	"crud-users/docs"
	"crud-users/models"
	"crud-users/routers"

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
	dsn := "root:@tcp(http://localhost:3306)/db_desafio?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate will create the necessary table. You can also use CreateTable to only create the table if it does not exist.
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate the database")
	}

	r := SetupRouter(db)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Run the application
	r.Run(":8080")
}
