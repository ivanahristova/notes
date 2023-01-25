package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"notes/backend/services/database"
	"notes/backend/utilities/token"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required,gte=3"`
	Password string `json:"password" binding:"required,gte=6"`
}

func GetCurrentUser(c *gin.Context) {
	userID, err := token.ExtractUserID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var user database.User
	user, err = database.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
	}
}

func GetUserNotesByID(c *gin.Context) {
	roleID, err := token.ExtractRoleID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var adminRoleID uint
	adminRoleID, err = database.GetRoleID("admin")

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	if roleID != adminRoleID {
		c.JSON(http.StatusForbidden, gin.H{"status": "fail", "data": "Unauthorized"})
		return
	}

	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var notes []database.Note
	notes, err = database.GetNotes(uint(userID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": notes})
	}
}

func Login(c *gin.Context) {
	var input LoginInput
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	var tkn string
	tkn, err = database.LoginUser(input.Username, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": "Incorrect username or password"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": tkn})
	}
}

func Signup(c *gin.Context) {
	var input SignupInput
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
		return
	}

	if err = database.AddUser(input.Email, input.Username, input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "data": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": "Registration successful"})
	}
}
