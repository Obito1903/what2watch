package groups

import (
	"db/pkg/utils"

	"github.com/gofiber/fiber/v3"
)

func RegisterGroupsRoute(app *fiber.App) {
	groups := app.Group("/groups")
	groups.Get("/", GetGroups)
	groups.Post("/", PostGroup)

	groups.Get("/:group_id", GetGroup)

	groups.Get("/:group_id/recommendations", GetGroupRecommendations)

	groups.Get("/:group_id/users", GetGroupUsers)
	groups.Put("/:group_id/users/:user_id", PutGroupUsers)
	groups.Delete("/:group_id/users/:user_id", DeleteGroupMovies)

	groups.Get("/:group_id/tastes", GetGroupTastes)

}

func GetGroups(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/groups")
}

func PostGroup(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/groups")
}

func GetGroup(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/groups/"+c.Params("group_id"))
}

func GetGroupRecommendations(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/groups/"+c.Params("group_id")+"/recommendations")
}

func GetGroupUsers(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/groups/"+c.Params("group_id")+"/users")
}

func PutGroupUsers(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/groups/"+c.Params("group_id")+"/users/"+c.Params("user_id"))
}

func DeleteGroupMovies(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/groups/"+c.Params("group_id")+"/users/"+c.Params("user_id"))
}

func GetGroupTastes(c fiber.Ctx) error {
	return utils.AuthProxyWrapper(c, utils.AppConfig.DBApiURL+"/groups/"+c.Params("group_id")+"/tastes")
}
