package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string //`gorm:"size:255;not null;unique" json:"email"`
	Username string //`gorm:"size:255;not null;unique" json:"username"`
	Password string //`gorm:"size:255;not null;" json:"password"`
}

func (u *User) PrepareGive() {
	u.Password = ""
}
