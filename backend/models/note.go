package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title       string `gorm:"size:255;not null;unique" json:"title"`
	Description string `gorm:"size:255;not null;" json:"description"`
	Status      string `gorm:"size:255;not null;" json:"status"`
	// UserId      string `gorm:"size:255;not null;unique" json:"userId"`
}
