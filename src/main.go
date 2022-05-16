package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/joho/godotenv"

	"api/src/helpers"
	"api/src/models"
	"api/src/router"
)

func init() {
	if envLoadError := godotenv.Load(); envLoadError != nil {
		log.Fatal("[ ERROR ] Failed to load .env file")
	}
}

func main() {
	var PORT string
	db := helpers.CreateDatabaseInstance()

	router := router.RegisterRoutes(db)

	if migrateError := db.AutoMigrate(&models.Film{}, &models.Crew{}, &models.Cast{}, &models.FilmRating{}, &models.FilmViewed{}, &models.FilmGenres{}, &models.User{}); migrateError != nil {
		log.Fatal("[ ERROR ] Couldn't migrate models!")
	}

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "9090"
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		Debug:          false,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete},
	})

	handler := c.Handler(router)

	fmt.Printf("[ OK ] Server is Started and Listening on port: %v", PORT)
	//helpers.CreateTmdbClient(db)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
