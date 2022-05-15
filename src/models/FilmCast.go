package models

import "gorm.io/gorm"

type Cast struct {
	gorm.Model
	Character   string
	Gender      int
	ID          uint
	Name        string
	Order       int
	ProfilePath string
}
