package database

import (
	"errors"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title       string `gorm:"size:255;not null;"`
	Description string `gorm:"size:255;not null;"`
	UserID      uint   `gorm:"size:255;not null;"`
}

func GetUserNotes(userId uint) ([]Note, error) {
	var notes []Note

	if err := database.Where("user_id = ?", userId).Find(&notes).Error; err != nil {
		return notes, errors.New("user not found")
	}

	return notes, nil
}

func AddNote(title, description string, userID uint) error {
	note := Note{
		Title:       html.EscapeString(strings.TrimSpace(title)),
		Description: html.EscapeString(strings.TrimSpace(description)),
		UserID:      userID,
	}

	return database.Create(&note).Error
}
