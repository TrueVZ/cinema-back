package models

import "gorm.io/gorm"

type Film struct {
	gorm.Model
	ID           uint
	Title        string
	IMDbID       string
	Genres       []FilmGenres `gorm:"many2many:films_genres;"`
	Cast         []Cast       `gorm:"many2many:film_cast;"`
	Crew         []Crew       `gorm:"many2many:film_crew;"`
	ReleaseDate  string
	Overview     string
	VoteAverage  float32
	Status       string
	PosterPath   string
	BackdropPath string
	Revenue      int64
	Runtime      int
	Adult        bool
	Budget       int64
}
