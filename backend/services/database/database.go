package database

import (
	"errors"
	"fmt"
	"html"
	"log"
	"os"
	"strings"

	"notes/backend/models"
	"notes/backend/utilities/token"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Could not load .env file")
	}

	dsn := os.Getenv("DATABASE_DNS")
	database, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println("Could not connect to database ")
		log.Fatal("connection error:", err)
	}

	database.AutoMigrate(&models.User{})

	fmt.Println("Database connection successful ")
}

func GetUserByID(uid uint) (models.User, error) {
	var user models.User

	if err := database.Omit("password").First(&user, uid).Error; err != nil {
		return user, errors.New("User not found")
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

	if err = verifyPassword(password, user.Password); err != nil {
		return "", err
	}

	tkn, err := token.Generate(user.ID)

	if err != nil {
		return "", err
	}

	return tkn, nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
}
