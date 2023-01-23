package database

import (
	"errors"
	"fmt"
	"html"
	"log"
	"notes/backend/utilities/token"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func Connect() {
	var err error

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Could not load .env file")
	}

	dsn := os.Getenv("DATABASE_DNS")

	database, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println("Could not connect to database ")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("Database connection successful ")
	}

	database.AutoMigrate(&User{})
}

func GetUserByID(uid uint) (User, error) {
	var u User

	if err := database.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found")
	}

	u.PrepareGive()

	return u, nil
}

func AddUser(email, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 15)

	if err != nil {
		return err
	}

	user := User{
		Email:    html.EscapeString(strings.TrimSpace(email)),
		Username: html.EscapeString(strings.TrimSpace(username)),
		Password: string(hashedPassword),
	}

	err = database.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func LoginUser(username, password string) (string, error) {
	var err error

	u := User{}

	err = database.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = verifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.Generate(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))
}
