package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string
	Viewed   []Film `gorm:"many2many:user_favorite;"`
}
