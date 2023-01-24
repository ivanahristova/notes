package database

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"notes/backend/models"
	"notes/backend/utilities/token"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalln("godotenv: could not load .env file")
	}

	dsn := os.Getenv("DATABASE_DNS")
	database, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalln("gorm: could not connect to database", err)
	}

	database.AutoMigrate(&models.User{})

	log.Println("Database connection successful")
}

func GetUserByID(uid uint) (models.User, error) {
	var user models.User

	if err := database.Omit("password").First(&user, uid).Error; err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}

func AddUser(email, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)

	if err != nil {
		return err
	}

	user := models.User{
		Email:    html.EscapeString(strings.TrimSpace(email)),
		Username: html.EscapeString(strings.TrimSpace(username)),
		Password: string(hashedPassword),
	}

	return database.Create(&user).Error
}

func LoginUser(username, password string) (string, error) {
	var user models.User
	var err error

	if err = database.Where("username = ?", username).Take(&user).Error; err != nil {
		return "", err
	}

	if err = verifyPassword(user.Password, password); err != nil {
		return "", err
	}

	tkn, err := token.Generate(user.ID, user.Admin)

	if err != nil {
		return "", err
	}

	return tkn, nil
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUserNotes(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error while parsing parameter %s", err.Error()),
		})
		return
	}

	var notes []models.Note
	if err := database.Where("user_id = ?", userId).Find(&notes).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while getting notes from database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": notes})
}
