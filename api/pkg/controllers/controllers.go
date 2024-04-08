package controllers

import (
	"context"
	"db/pkg/config"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GoogleLogin(c *fiber.Ctx) error {

	url := config.AppConfig.LoginConfig.AuthCodeURL("randomstate")

	log.Info(url)
	c.Status(fiber.StatusSeeOther)
	c.Redirect(url)
	return c.JSON(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "randomstate" {
		return c.SendString("States don't Match!!")
	}

	code := c.Query("code")

	googlecon := config.KeycloakConfig()
	log.Info("Code: ", code)

	token, err := googlecon.Exchange(context.Background(), code)
	if err != nil {
		log.Error(err)
		return c.SendString("Code-Token Exchange Failed")
	}
	log.Info("Token: ", token.AccessToken)
	req, err := http.NewRequest("GET", "http://auth.localhost/realms/what2watch/protocol/openid-connect/userinfo", nil)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return c.SendString("User Data Fetch Failed")
	}

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.SendString("JSON Parsing Failed")
	}

	return c.SendString(string(userData))

}
