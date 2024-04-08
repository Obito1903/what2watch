package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

type Config struct {
	LoginConfig oauth2.Config
}

var AppConfig Config

func KeycloakConfig() oauth2.Config {
	err := godotenv.Load(".env")
	log.Println("Loading .env file")
	log.Println(os.Getenv("KEYCLOAK_CLIENT_ID"))

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
