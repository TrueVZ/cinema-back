package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	controller "api/src/controllers"
	middleware "api/src/middlewares"
)

func RegisterRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	FoodController := controller.FoodController{}

	router.HandleFunc("/food/all", middleware.CheckAuth(FoodController.GetAllFoodItems(db))).Methods(http.MethodGet).Headers()
	router.HandleFunc("/food", middleware.CheckAuth(FoodController.AddNewFoodItem(db))).Methods(http.MethodPost)
	router.HandleFunc("/food/{name}", middleware.CheckAuth(FoodController.GetSingleFoodItem(db))).Methods(http.MethodGet)
	router.HandleFunc("/food/{id}", middleware.CheckAuth(FoodController.DeleteSingleFoodItem(db))).Methods(http.MethodDelete)

	UserController := controller.UserController{}

	router.HandleFunc("/auth/login", UserController.LoginUser(db)).Methods(http.MethodPost)
	router.HandleFunc("/auth/signup", UserController.SignupUser(db)).Methods(http.MethodPost)

	FilmController := controller.FilmController{}

	router.HandleFunc("/film/all", middleware.CheckAuth(FilmController.GetAllFilms(db))).Methods(http.MethodGet)
	router.HandleFunc("/film/{id}", middleware.CheckAuth(FilmController.GetFilmById(db))).Methods(http.MethodGet)
	router.HandleFunc("/film/{id}/addFavorite", middleware.CheckAuth(FilmController.AddFilmToFavorite(db))).Methods(http.MethodPost)
	router.HandleFunc("/getFavorites", middleware.CheckAuth(FilmController.GetFavoritesFilms(db))).Methods(http.MethodGet)
	router.HandleFunc("/film/{id}/removeFavorite", middleware.CheckAuth(FilmController.RemoveFromFavorite(db))).Methods(http.MethodDelete)

	return router
}
