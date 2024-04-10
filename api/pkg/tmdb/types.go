package tmdb

import (
	tmdbApi "github.com/cyruzin/golang-tmdb"
)

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Movie struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Poster      string  `json:"poster"`
	ReleaseDate string  `json:"release_date"`
	Genres      []Genre `json:"genres"`
}

func movieFromMoviePopularResults(tmdbMovies *tmdbApi.MoviePopularResults) []Movie {

	var movies []Movie

	for _, movie := range tmdbMovies.Results {
		genres := []Genre{}
		for _, genre := range movie.Genres {
			genres = append(genres, Genre{
				ID:   genre.ID,
				Name: genre.Name,
			})
		}

		movies = append(movies, Movie{
			ID:          movie.ID,
			Title:       movie.Title,
			Poster:      "https://image.tmdb.org/t/p/w300_and_h450_bestv2" + movie.PosterPath,
			ReleaseDate: movie.ReleaseDate,
			Genres:      genres,
		})
	}

	return movies
}
