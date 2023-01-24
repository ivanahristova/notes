package controllers

import (
	"net/http"
	"notes/backend/models"
	"notes/backend/services/database"
	"notes/backend/utilities/token"

	"github.com/gin-gonic/gin"
)

type NoteInput struct {
	Title       string `json:"title" binding:"required,gte=1"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required,gte=1"`
}

// func AddNote(c *gin.Context) {
// 	var input NoteInput
// 	var err error

// 	if err = c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err = database.AddNote(input.Title, input.Description, input.Status)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "general error"})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{"message": "success"})
// 	}
// }

func GetUserNotes(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var notes []models.Note

	notes, err = database.GetUserNotes(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": notes})
	}
}
