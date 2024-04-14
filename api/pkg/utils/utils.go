package utils

import (
	"context"
	"db/pkg/dbapi"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"github.com/joho/godotenv"
)

type ApiError struct {
	Message string `json:"message"`
}

type UserInfo struct {
	Sub               string `json:"sub"`
	Email             string `json:"email"`
	PreferredUsername string `json:"preferred_username"`
}

func CheckAuth(c fiber.Ctx) (int, error) {
	user, err := CheckAuthKeycloak(c)
	if err != nil {
		return 0, err
	}
	ctx := context.Background()
	userResp, err := AppConfig.DBClient.GetUserInfosByMailUsersEmailBymailGetWithResponse(ctx, user.Email)
	if err != nil {
		return 0, err
	}
	if userResp.StatusCode() != 200 {
		return 0, nil
	}
	return userResp.JSON200.UserId, nil
}

func AuthProxyWrapper(c fiber.Ctx, addr string) error {
	_, err := CheckAuth(c)
	if err != nil {
		return c.Status(401).JSON(ApiError{Message: "Unauthorized"})
	}
	return proxy.Forward(addr)(c)
}

func CheckAuthKeycloak(c fiber.Ctx) (*UserInfo, error) {
	// token := c.Cookies("session-token")
	// fmt.Println(token)
	// return &UserInfo{
	// 	Sub:               "123",
	// 	Email:             "test@example.com",
	// 	PreferredUsername: "test",
	// }, nil
	req, err := http.NewRequest("GET", AppConfig.AuthURL, nil)
	if err != nil {
		log.Error(fmt.Sprintf("Error creating request:\n%v", err))
		return nil, errors.New("error creating request")
	}
	if c.Get("Authorization") == "" {
		log.Error("No Authorization header")
		return nil, errors.New("no Authorization header")
	}
	token := c.Get("Authorization")[7:]
	req.Header.Set("Authorization", "Bearer "+token)
	// req.Header.Add("Authorization", c.Get("Authorization"))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error(fmt.Sprintf("Error sending request:\n%v", err))
		return nil, errors.New("error sending request")
	}
	if res.StatusCode != 200 {
		log.Error(fmt.Sprintf("Error validating token: %v", res.Status))
		body, _ := io.ReadAll(res.Body)
		fmt.Println(body)
		return nil, errors.New("error validating token")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(fmt.Sprintf("Error reading response:\n%v", err))
		return nil, errors.New("error reading response")
	}
	userInfo := new(UserInfo)
	if err := json.Unmarshal(body, userInfo); err != nil {
		log.Error(fmt.Sprintf("Error parsing response:\n%v", err))
		return nil, errors.New("error parsing response")
	}
	return userInfo, nil
}

type Config struct {
	AuthURL    string
	DBApiURL   string
	EngineUrl  string
	DBClient   *dbapi.ClientWithResponses
	TmdbClient *tmdb.Client
}

var AppConfig Config

func KeycloakConfig() string {
	godotenv.Load(".env")

	AppConfig.AuthURL = os.Getenv("AUTH_URL")

	return AppConfig.AuthURL
}

func DBApiConfig() string {
	godotenv.Load(".env")

	AppConfig.DBApiURL = os.Getenv("DB_API_URL")
	client, err := dbapi.NewClientWithResponses(AppConfig.DBApiURL)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
	AppConfig.DBClient = client
	return AppConfig.DBApiURL
}

func TmdbConfig() *tmdb.Client {
	godotenv.Load(".env")
	var err error

	AppConfig.TmdbClient, err = tmdb.Init(os.Getenv("TMDB_API_KEY"))
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
	return AppConfig.TmdbClient
}

func EngineConfig() string {
	godotenv.Load(".env")
	AppConfig.EngineUrl = os.Getenv("ENGINE_URL")
	return AppConfig.EngineUrl
}
