package utils

import (
	"db/pkg/dbapi"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type ApiError struct {
	Message string `json:"message"`
}

type UserInfo struct {
	Sub               string `json:"sub"`
	Email             string `json:"email"`
	PreferredUsername string `json:"preferred_username"`
}

func CheckAuth(c fiber.Ctx) (*UserInfo, error) {
	return &UserInfo{
		Sub:               "123",
		Email:             "test@example.com",
		PreferredUsername: "test",
	}, nil
	// req, err := http.NewRequest("GET", "http://auth.localhost/realms/what2watch/protocol/openid-connect/userinfo", nil)
	// if err != nil {
	// 	log.Error(fmt.Sprintf("Error creating request:\n%v", err))
	// 	return nil, errors.New("error creating request")
	// }
	// req.Header.Add("Authorization", c.Get("Authorization"))
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	log.Error(fmt.Sprintf("Error sending request:\n%v", err))
	// 	return nil, errors.New("error sending request")
	// }
	// if res.StatusCode != 200 {
	// 	log.Error(fmt.Sprintf("Error validating token: %v", res.Status))
	// 	return nil, errors.New("error validating token")
	// }
	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Error(fmt.Sprintf("Error reading response:\n%v", err))
	// 	return nil, errors.New("error reading response")
	// }
	// userInfo := new(UserInfo)
	// if err := json.Unmarshal(body, userInfo); err != nil {
	// 	log.Error(fmt.Sprintf("Error parsing response:\n%v", err))
	// 	return nil, errors.New("error parsing response")
	// }
	// return userInfo, nil
}

type Config struct {
	LoginConfig oauth2.Config
	DBApiURL    string
	DBClient    *dbapi.ClientWithResponses
}

var AppConfig Config

func KeycloakConfig() oauth2.Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.LoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:3000/callback",
		ClientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		ClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   os.Getenv("KEYCLOAK_AUTH_URL"),
			TokenURL:  os.Getenv("KEYCLOAK_TOKEN_URL"),
			AuthStyle: oauth2.AuthStyleAutoDetect,
		},
	}

	return AppConfig.LoginConfig
}

func DBApiConfig() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	AppConfig.DBApiURL = os.Getenv("DB_API_URL")
	client, err := dbapi.NewClientWithResponses(AppConfig.DBApiURL)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
	AppConfig.DBClient = client
	return AppConfig.DBApiURL
}
