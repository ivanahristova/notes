package controllers

import (
	"net/http"
	"notes/backend/services/database"
	"notes/backend/utilities/token"

	"github.com/gin-gonic/gin"
)

type NoteInput struct {
	Title       string `json:"title" binding:"required,gte=1"`
	Description string `json:"description"`
}

func GetUserNotes(c *gin.Context) {
	userID, err := token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var notes []database.Note
	notes, err = database.GetUserNotes(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": notes})
	}
}

func AddNote(c *gin.Context) {
	var input NoteInput
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var userID uint
	userID, err = token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	err = database.AddNote(input.Title, input.Description, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": "Note added successfully"})
	}
}
