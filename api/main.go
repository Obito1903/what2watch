package main

import (
	"db/pkg/users"
	"db/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func checkAuth(token string) bool {
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

	utils.KeycloakConfig()
	utils.DBApiConfig()

	users.RegisterUsersRoutes(app)

	log.SetLevel(log.LevelDebug)
	log.Fatal(app.Listen(":3000"))
}
