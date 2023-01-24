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

func GetUserNotes(c *gin.Context) {
	userID, err := token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var notes []models.Note
	notes, err = database.GetUserNotes(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": notes})
	}
}
