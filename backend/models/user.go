package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique"`
	Username string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null;"`
	RoleID   uint   `gorm:"not null;"`
}
