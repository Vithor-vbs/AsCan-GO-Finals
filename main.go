package main

import (
	"log"

	"github.com/Vithor-vbs/API-Requests-Example/controllers"
	"github.com/Vithor-vbs/API-Requests-Example/initializers"
	"github.com/Vithor-vbs/API-Requests-Example/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
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

	router.Run()
}
