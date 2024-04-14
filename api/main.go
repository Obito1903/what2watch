package main

import (
	"db/pkg/groups"
	"db/pkg/tmdb"
	"db/pkg/users"
	"db/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func _checkAuth(token string) bool {
	// Keycloak API call to validate token using headers
	req, err := http.NewRequest("GET", "http://auth.localhost/realms/what2watch/protocol/openid-connect/userinfo", nil)
	if err != nil {
		log.Error(err)
		return false
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Error(err)
		return false
	}
	if res.StatusCode != 200 {
		return false
	}
	return true
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:        "http://localhost:5173, http://what2watch.localhost",
		AllowHeaders:        "Origin, Content-Type, Accept, Authorization",
		AllowCredentials:    true,
		AllowPrivateNetwork: true,
	}))

	utils.KeycloakConfig()
	utils.DBApiConfig()
	utils.TmdbConfig()

	users.RegisterUsersRoutes(app)
	tmdb.RegisterTmdbRoutes(app)
	groups.RegisterGroupsRoute(app)

	log.SetLevel(log.LevelDebug)
	log.Fatal(app.Listen(":3000"))
}
