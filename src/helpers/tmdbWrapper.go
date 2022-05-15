package helpers

import (
	"api/src/models"
	"fmt"
	"github.com/cyruzin/golang-tmdb"
	"gorm.io/gorm"
)

func CreateTmdbClient(db *gorm.DB) {
	tmdbClient, err := tmdb.Init("e1bb876eaa8453b25db3a03c61947090")
	if err != nil {
		fmt.Println(err)
	}
	options := map[string]string{
		"language":           "ru-RU",
		"append_to_response": "credits,images",
	}
	popularMovies, err := tmdbClient.GetMoviePopular(options)
	if err != nil {
		fmt.Println(err)
	}
	for _, popularMovie := range popularMovies.Results {
		movie, err := tmdbClient.GetMovieDetails(int(popularMovie.ID), options)
		if err != nil {
			fmt.Println(err)
		}
		var genresInDb []models.FilmGenres
		var crewsInDb []models.Crew
		var castsInDb []models.Cast

		for _, genre := range movie.Genres {
			genreInDb := models.FilmGenres{
				ID:   uint(genre.ID),
				Name: genre.Name,
			}
			genresInDb = append(genresInDb, genreInDb)
		}

		for _, crew := range movie.Credits.Crew {
			CrewInDb := models.Crew{
				ID:          uint(crew.ID),
				Name:        crew.Name,
				Department:  crew.Department,
				Gender:      crew.Gender,
				Job:         crew.Job,
				Popularity:  crew.Popularity,
				ProfilePath: crew.ProfilePath,
			}
			if result := db.FirstOrCreate(&CrewInDb); result.Error != nil {
				fmt.Println(result.Error.Error())
			}
			crewsInDb = append(crewsInDb, CrewInDb)
		}
		for _, cast := range movie.Credits.Cast {
			CastInDb := models.Cast{
				ID:          uint(cast.ID),
				Character:   cast.Character,
				Gender:      cast.Gender,
				Name:        cast.Name,
				Order:       cast.Order,
				ProfilePath: cast.ProfilePath,
			}
			if result := db.FirstOrCreate(&CastInDb); result.Error != nil {
				fmt.Println(result.Error.Error())
			}
			castsInDb = append(castsInDb, CastInDb)
		}

		movieInDb := models.Film{
			ID:           uint(movie.ID),
			Title:        movie.Title,
			IMDbID:       movie.IMDbID,
			Genres:       genresInDb,
			Cast:         castsInDb,
			Crew:         crewsInDb,
			ReleaseDate:  movie.ReleaseDate,
			Overview:     movie.Overview,
			VoteAverage:  movie.VoteAverage,
			Status:       movie.Status,
			PosterPath:   movie.PosterPath,
			BackdropPath: movie.BackdropPath,
			Revenue:      movie.Revenue,
			Runtime:      movie.Runtime,
			Adult:        movie.Adult,
			Budget:       movie.Budget,
		}
		if result := db.FirstOrCreate(&movieInDb); result.Error != nil {
			fmt.Println(result.Error.Error())
		}
	}
}
