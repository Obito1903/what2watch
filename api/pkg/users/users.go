package users

import (
	"context"
	"db/pkg/dbapi"
	"db/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/proxy"
)

func RegisterUsersRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/me", GetMe)
	users.Post("/me", PostMe)
	users.Delete("/me", DeleteMe)

	users.Get("/movies", GetMeMovies)
	users.Post("/movies/:movie_id", PostMeMovies)
	users.Delete("/movies/:movie_id", DeleteMeMovies)

	users.Get("/:user_id/movies", GetUserMovies)

	users.Get("/recommendations", GetMeRecommenations)
	users.Get("/:user_id/recommendations/", GetUserRecommendations)

	users.Get("/tastes", GetMeTastes)
	users.Put("/tastes/:genre_id", PostMeTastes)
	users.Delete("/tastes/:genre_id", DeleteMeTastes)

}

func GetMe(c fiber.Ctx) error {
	user, err := utils.CheckAuthKeycloak(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})

	}
	ctx := context.Background()
	// Check if user is in the database
	userResp, err := utils.AppConfig.DBClient.GetUserInfosByMailUsersEmailBymailGetWithResponse(ctx, user.Email)
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(utils.ApiError{Message: "Error getting user"})
	}
	if userResp.StatusCode() == 404 {
		// If not, add user to the database
		resp, err := utils.AppConfig.DBClient.CreateUserUsersPostWithResponse(ctx, dbapi.UserPostCreateRequest{
			Mail: user.Email,
			Name: user.PreferredUsername,
		})
		if err != nil {
			log.Error(err)
			return c.Status(500).JSON(utils.ApiError{Message: "Error Creating user"})
		}
		if resp.StatusCode() != 200 {
			return c.Status(resp.StatusCode()).JSON(utils.ApiError{Message: "Error Creating user"})
		}
		userResp, err := utils.AppConfig.DBClient.GetUserInfosByMailUsersEmailBymailGetWithResponse(ctx, user.Email)
		if err != nil {
			log.Error(err)
			return c.Status(500).JSON(utils.ApiError{Message: "Error getting user"})
		}
		if userResp.StatusCode() != 200 {
			return c.Status(userResp.StatusCode()).JSON(utils.ApiError{Message: "Error getting user"})
		}
		return c.JSON(userResp.JSON200)
	}

	if userResp.StatusCode() != 200 {
		return c.Status(userResp.StatusCode()).JSON(utils.ApiError{Message: "Error getting user"})
	}

	return c.JSON(userResp.JSON200)
}

func PostMe(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d", utils.AppConfig.DBApiURL, userId))(c)
}

func DeleteMe(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d", utils.AppConfig.DBApiURL, userId))(c)
}

func GetMeMovies(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d/movies", utils.AppConfig.DBApiURL, userId))(c)
}

func PostMeMovies(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d/movies/%s", utils.AppConfig.DBApiURL, userId, c.Params("movie_id")))(c)
}

func DeleteMeMovies(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d/movies/%s", utils.AppConfig.DBApiURL, userId, c.Params("movie_id")))(c)
}

func GetUserMovies(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/users/"+c.Params("user_id")+"/movies")
}

func GetMeRecommenations(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d/recommendations", utils.AppConfig.DBApiURL, userId))(c)
}

func GetUserRecommendations(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/users/recommendations")
}

func GetMeTastes(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d/tastes", utils.AppConfig.DBApiURL, userId))(c)
}

func PostMeTastes(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d/tastes/%s", utils.AppConfig.DBApiURL, userId, c.Params("genre_id")))(c)
}

func DeleteMeTastes(c fiber.Ctx) error {
	userId, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(fmt.Sprintf("%s/users/%d/tastes/%s", utils.AppConfig.DBApiURL, userId, c.Params("genre_id")))(c)
}
