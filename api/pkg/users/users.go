package users

import (
	"context"
	"db/pkg/dbapi"
	"db/pkg/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func RegisterUsersRoutes(app *fiber.App) {
	users := app.Group("/users")
	users.Get("/me", GetMe)
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
		resp, err := utils.AppConfig.DBClient.CreateUserUsersPostWithResponse(ctx, &dbapi.CreateUserUsersPostParams{
			Name: user.PreferredUsername,
			Mail: user.Email,
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

func GetUserRecommenations(c fiber.Ctx) error {
	userID, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	rec, err := utils.AppConfig.DBClient.GetRecommendationsUsersUserIdRecommendationsGetWithResponse(context.Background(), userID)
	if err != nil {
		log.Error(err)
		return c.Status(500).JSON(utils.ApiError{Message: "Error getting recommendations"})
	}
	if rec.StatusCode() != 200 {
		return c.Status(rec.StatusCode()).SendString(string(rec.Body))
	}
	return c.JSON(rec.JSON200)
}
