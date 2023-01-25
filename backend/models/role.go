package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Code string `gorm:"size:255;not null;unique"`
}
