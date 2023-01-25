package database

import (
	"errors"
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
		log.Fatalln("godotenv: could not load .env file")
	}

	dsn := os.Getenv("DATABASE_DNS")
	database, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalln("gorm: could not connect to database", err)
	}

	if err = database.AutoMigrate(&models.User{}, &models.Note{}, &models.Role{}); err != nil {
		log.Fatalln("gorm: could not run auto migration")
	}

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

	var userRoleID uint
	userRoleID, err = GetRoleID("user")

	if err != nil {
		return err
	}

	user := models.User{
		Email:    html.EscapeString(strings.TrimSpace(email)),
		Username: html.EscapeString(strings.TrimSpace(username)),
		Password: string(hashedPassword),
		RoleID:   userRoleID,
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

	var tkn string
	tkn, err = token.Generate(user.ID, user.RoleID)

	if err != nil {
		return "", err
	}

	return tkn, nil
}

func GetUserNotes(userId uint) ([]models.Note, error) {
	var notes []models.Note

	if err := database.Where("user_id = ?", userId).Find(&notes).Error; err != nil {
		return notes, errors.New("user not found")
	}

	return notes, nil
}

func GetRoleID(code string) (uint, error) {
	var role models.Role

	if err := database.Where("code = ?", code).Find(&role).Error; err != nil {
		return 0, err
	}

	return role.ID, nil
}

func AddNote(title, description string, userID uint) error {
	note := models.Note{
		Title:       html.EscapeString(strings.TrimSpace(title)),
		Description: html.EscapeString(strings.TrimSpace(description)),
		UserID:      userID,
	}

	return database.Create(&note).Error
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
