package main

import (
	"db/pkg/config"
	"db/pkg/controllers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

	config.KeycloakConfig()

	app.Post("/login", controllers.GoogleLogin)
	app.Get("/callback", controllers.GoogleCallback)
	app.Get("/auth", func(c *fiber.Ctx) error {
		token := c.Query("token")
		if checkAuth(token) {
			return c.SendString("Authenticated")
		}
		return c.SendString("Not Authenticated")
	})

	log.SetLevel(log.LevelDebug)
	log.Fatal(app.Listen(":3000"))
}
