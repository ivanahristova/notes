package controllers

import (
	"net/http"

	"notes/backend/models"
	"notes/backend/services/database"
	"notes/backend/utilities/token"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user, err = database.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
	}
}

func Login(c *gin.Context) {
	var input LoginInput
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tkn string
	tkn, err = database.LoginUser(input.Username, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect username or password"})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": tkn})
	}
}

func Signup(c *gin.Context) {
	var input SignupInput
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = database.AddUser(input.Email, input.Username, input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
	}
}
