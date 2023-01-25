package main

import (
	"notes/backend/controllers"
	"notes/backend/middlewares"
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

	notes := api.Group("/notes")
	notes.Use(middlewares.JwtAuthMiddleware())
	notes.GET("/", controllers.GetUserNotes)
	notes.POST("/new", controllers.AddNote)

	admin := api.Group("/admin")
	admin.Use(middlewares.JwtAuthMiddleware())
	admin.GET("/users/:id", controllers.GetUserNotesByID)

	router.Run(":8080")
}
