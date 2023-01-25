package database

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Code string `gorm:"size:255;not null;unique"`
}

func GetRoleID(code string) (uint, error) {
	var role Role

	if err := database.Where("code = ?", code).Find(&role).Error; err != nil {
		return 0, err
	}

	return role.ID, nil
}
