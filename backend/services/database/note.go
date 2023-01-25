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

func GetNote(noteID uint) (Note, error) {
	var note Note

	if err := database.First(&note, noteID).Error; err != nil {
		return note, errors.New("note not found")
	}

	return note, nil
}

func GetNotes(userID uint) ([]Note, error) {
	var notes []Note

	if err := database.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
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

	if err := database.Create(&note).Error; err != nil {
		return errors.New("could not add note")
	}

	return nil
}

func UpdateNote(id uint, title, description string) error {
	var note Note
	var err error

	note, err = GetNote(id)

	if err != nil {
		return err
	}

	note.Title = title
	note.Description = description

	if err = database.Save(&note).Error; err != nil {
		return errors.New("could not update note")
	}

	return nil
}

func DeleteNote(id uint) error {
	if err := database.Delete(&Note{}, id).Error; err != nil {
		return errors.New("could not delete note")
	}

	return nil
}
