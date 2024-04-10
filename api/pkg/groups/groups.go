package groups

import (
	"db/pkg/dbapi"
	"db/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

func RegisterGroupsRoute(app *fiber.App) {
	groups := app.Group("/groups")
	groups.Get("/", GetGroups)
	groups.Post("/", PostGroup)
	groups.Get("/:id", GetGroup)
}

func GetGroups(c fiber.Ctx) error {
	userID, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	groups, err := utils.AppConfig.DBClient.GetUserGroupsUserUserIdGroupsGet(c.Context(), userID)
	if err != nil {
		return c.Status(500).JSON(utils.ApiError{Message: "Error getting groups"})
	}
	return c.JSON(groups)
}

func GetGroup(c fiber.Ctx) error {
	return nil
}

func PostGroup(c fiber.Ctx) error {
	userID, err := utils.CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(utils.ApiError{Message: "Unauthorized"})
	}
	group := dbapi.CreateGroupGroupsPostParams{
		GpName: c.FormValue("name"),
	}
	body := c.Body()
	group.UserID = userID
	groupResp, err := utils.AppConfig.DBClient.CreateGroupGroupsPostWithResponse(c.Context())
	if err != nil {
		return c.Status(500).JSON(utils.ApiError{Message: "Error creating group"})
	}
	return c.JSON(groupResp)
}
