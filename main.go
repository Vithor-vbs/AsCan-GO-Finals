package main

import (
	"log"

	"github.com/Vithor-vbs/API-Requests-Example/controllers"
	"github.com/Vithor-vbs/API-Requests-Example/initializers"
	"github.com/Vithor-vbs/API-Requests-Example/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title Swagger Ascan Finals API
// @versoin 1.0
// @description This is a demo server
// @BasePath /api/v1
// @localhost:8080
func main() {

	// Set swagger meta information
	docs.SwaggerInfo.Title = "Swagger Ascan Finals API"
	docs.SwaggerInfo.Description = "Ascan finals api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	initializers.ConnectToDB()

	if err := initializers.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to autoMigrate database")
	}

	router.POST("/user/createWithArray", controllers.CreateUserByArr)

	router.GET("/user", controllers.GetUsers)
	router.GET("/user/:username", controllers.FindUser)
	router.POST("/user", controllers.CreateUser)
	router.PATCH("/user/:id", controllers.UpdateUser)
	router.PUT("/user/:id", controllers.PutUser)
	router.DELETE("/user/:id", controllers.DeleteUser)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
