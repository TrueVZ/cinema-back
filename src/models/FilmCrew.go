package models

import "gorm.io/gorm"

type Crew struct {
	gorm.Model
	Department  string
	Gender      int
	ID          uint
	Job         string
	Name        string
	Popularity  float32
	ProfilePath string
}
