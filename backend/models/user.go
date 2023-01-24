package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"Email"`
	Username string `gorm:"size:255;not null;unique" json:"Username"`
	Password string `gorm:"size:255;not null;" json:"Password"`
	RoleID   uint   `gorm:"not null" json:"RoleID"`
}
