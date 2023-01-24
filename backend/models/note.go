package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title       string `gorm:"size:255;not null;"`
	Description string `gorm:"size:255;not null;"`
	UserID      uint   `gorm:"size:255;not null;"`
}
