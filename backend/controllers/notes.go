package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"notes/backend/services/database"
	"notes/backend/utilities/token"
)

type NoteInput struct {
	Title       string `json:"title" binding:"required,gte=1"`
	Description string `json:"description"`
}

func Index(c *gin.Context) {
	userID, err := token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var notes []database.Note
	notes, err = database.GetNotes(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": notes})
	}
}

func Create(c *gin.Context) {
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

func Show(c *gin.Context) {
	noteID, err := strconv.ParseUint(c.Param("noteID"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var note database.Note
	note, err = database.GetNote(uint(noteID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": note})
}

func Update(c *gin.Context) {
	// TODO...
}

func Destroy(c *gin.Context) {
	// TODO...
}
