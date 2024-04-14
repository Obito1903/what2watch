package tmdb

import (
	"db/pkg/utils"
	"strconv"

	tmdbApi "github.com/cyruzin/golang-tmdb"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

var tmdbClient *tmdbApi.Client

func RegisterTmdbRoutes(app *fiber.App) {
	tmdbClient = utils.AppConfig.TmdbClient
	tmdb := app.Group("/tmdb")
	tmdb.Get("/movies/toprated", GetMoviesTopRated)
	tmdb.Get("/movies/popular", GetMoviesPopular)
	tmdb.Get("/movies/:movie_id/details", GetMovieDetails)
	tmdb.Get("/search", GetSearchMovie)
	tmdb.Get("/genres", GetGenres)
}

func GetMoviesTopRated(c fiber.Ctx) error {
	res, err := tmdbClient.GetMovieTopRated(map[string]string{"language": "en-US"})
	if err != nil {
		return c.Status(500).JSON(utils.ApiError{Message: "Error getting movies"})
	}

	movies := movieFromMoviePopularResults(res.MoviePopularResults)

	return c.JSON(movies)
}

func GetMoviesPopular(c fiber.Ctx) error {
	res, err := tmdbClient.GetMoviePopular(map[string]string{"language": "en-US"})
	if err != nil {
		return c.Status(500).JSON(utils.ApiError{Message: "Error getting movies"})
	}

	movies := movieFromMoviePopularResults(res.MoviePopularResults)

	return c.JSON(movies)
}

func GetSearchMovie(c fiber.Ctx) error {
	return proxy.Forward("https://api.themoviedb.org/3/search/movie")(c)
}

func GetMovieDetails(c fiber.Ctx) error {
	movieID, err := strconv.Atoi(c.Params("movie_id"))
	if err != nil {
		return c.Status(400).JSON(utils.ApiError{Message: "Invalid movie ID"})
	}
	res, err := tmdbClient.GetMovieDetails(movieID, map[string]string{"language": "en-US"})
	if err != nil {
		return c.Status(500).JSON(utils.ApiError{Message: "Error getting movie details"})
	}

	genres := []Genre{}

	for _, genre := range res.Genres {
		genres = append(genres, Genre{
			ID:   genre.ID,
			Name: genre.Name,
		})
	}

	movie := MovieDetails{
		ID:          res.ID,
		Title:       res.Title,
		Genres:      genres,
		Overview:    res.Overview,
		Popularity:  res.Popularity,
		ReleaseDate: res.ReleaseDate,
	}

	return c.JSON(movie)
}

func GetGenres(c fiber.Ctx) error {
    res, err := tmdbClient.GetGenreMovieList(map[string]string{"language": "en-US"})
    if err != nil {
        return c.Status(500).JSON(utils.ApiError{Message: "Error getting genres"})
    }

    genres := []Genre{}
    for _, genre := range res.Genres {
        genres = append(genres, Genre{
            ID:   genre.ID,
            Name: genre.Name,
        })
    }

    return c.JSON(genres)
}
