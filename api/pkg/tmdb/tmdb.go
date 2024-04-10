package tmdb

import (
	"db/pkg/utils"

	tmdbApi "github.com/cyruzin/golang-tmdb"
	"github.com/gofiber/fiber/v3"
)

var tmdbClient *tmdbApi.Client

func RegisterTmdbRoutes(app *fiber.App) {
	tmdbClient = utils.AppConfig.TmdbClient
	tmdb := app.Group("/tmdb")
	tmdb.Get("/movies/toprated", GetMoviesTopRated)
	tmdb.Get("/movies/popular", GetMoviesPopular)
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
