package main

import (
	"notes/backend/controllers"
	"notes/backend/services/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	api := router.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/signup", controllers.Signup)
	auth.POST("/login", controllers.Login)

	router.Run(":8080")
}
