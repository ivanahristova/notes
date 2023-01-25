package main

import (
	"github.com/gin-gonic/gin"

	"notes/backend/controllers"
	"notes/backend/middlewares"
	"notes/backend/services/database"
)

func main() {
	database.Connect()

	router := gin.Default()

	api := router.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/signup", controllers.Signup)
	auth.POST("/login", controllers.Login)

	notes := api.Group("/notes")
	notes.Use(middlewares.JwtAuthMiddleware())
	notes.GET("/", controllers.Index)
	notes.POST("/", controllers.Create)
	notes.GET("/:id", controllers.Show)
	notes.PUT("/:id", controllers.Update)
	notes.DELETE("/:id", controllers.Destroy)

	admin := api.Group("/admin")
	admin.Use(middlewares.JwtAuthMiddleware())
	admin.GET("/users/:id", controllers.GetUserNotesByID)

	router.Run(":8080")
}
