package models

import "gorm.io/gorm"

type FilmGenres struct {
	gorm.Model
	ID   uint
	Name string
}
