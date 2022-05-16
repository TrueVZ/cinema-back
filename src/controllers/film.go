package controller

import (
	"api/src/helpers"
	"api/src/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type FilmController struct{}

type ResponseFilm struct {
	Movie models.Film
}

var error = helpers.CustomError{}

func (film FilmController) GetAllFilms(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		var FilmItems []models.Film

		if results := db.Preload("Genres").Preload("Crew").Preload("Cast").Find(&FilmItems); results.Error != nil {
			error.ApiError(w, http.StatusInternalServerError, "Failed To Fetch Films Items from database!")
			return
		}

		helpers.RespondWithJSON(w, FilmItems)
	}
}

func (film FilmController) GetFilmById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		FilmItem := models.Film{}

		if results := db.Preload("Genres").Preload("Crew").Preload("Cast").Where("id = ?", params["id"]).First(&FilmItem); results.Error != nil || results.RowsAffected < 1 {
			error.ApiError(w, http.StatusNotFound, "Didn't Find film item with id = "+params["id"])
			return
		}

		helpers.RespondWithJSON(w,
			ResponseFilm{Movie: FilmItem},
		)
	}
}

func (FilmController) AddFilmToFavorite(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		authHeader := r.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")[1]
		claims, err := helpers.VerifyJwtToken(token)
		if err != nil {
			error.ApiError(w, http.StatusForbidden, err.Error())
			return
		}

		userId := claims.Id
		filmId, _ := strconv.ParseUint(params["id"], 10, 32)

		user := models.User{}
		film := models.Film{}

		db.First(&film, filmId)
		db.First(&user, userId)
		db.Model(&user).Association("Viewed").Append(&film)
		db.Preload("Viewed").First(&user, userId)
		helpers.RespondWithJSON(w, user.Viewed)
	}
}

func (FilmController) GetFavoritesFilms(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")[1]
		claims, err := helpers.VerifyJwtToken(token)
		if err != nil {
			error.ApiError(w, http.StatusForbidden, err.Error())
		}

		user := models.User{}
		userId := claims.Id

		db.Preload("Viewed").First(&user, userId)
		helpers.RespondWithJSON(w, user.Viewed)
	}
}

func (FilmController) RemoveFromFavorite(db *gorm.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		authHeader := request.Header.Get("Authorization")
		token := strings.Split(authHeader, " ")[1]
		claims, err := helpers.VerifyJwtToken(token)
		if err != nil {
			error.ApiError(writer, http.StatusForbidden, err.Error())
		}
		filmId, _ := strconv.ParseUint(params["id"], 10, 32)
		user := models.User{}
		film := models.Film{}
		userId := claims.Id
		db.First(&film, filmId)
		db.Preload("Viewed").First(&user, userId)
		result := db.Model(&user).Association("Viewed").Delete(&film)
		if result != nil {
			return
		}
		helpers.RespondWithJSON(writer, user.Viewed)
	}
}
