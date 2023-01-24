package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserId      string `gorm:"size:255;not null;unique" json:"userId"`
	Title       string `gorm:"size:255;not null;unique" json:"title"`
	Description string `gorm:"size:255;not null;" json:"description"`
	Date        string `gorm:"size:255;not null;" json:"date"`
	Status      string `gorm:"size:255;not null;" json:"status"`
	PriorityId  string `gorm:"size:255;not null;" json:"priorityId"`
}
