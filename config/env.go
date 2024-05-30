package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var BasePath string
var Port string

var Algorithm string
var JwtExpireMins int
var Secret []byte
var WebApp string
var ApiKey string

var DbDriver string
var DbUrl string

func LoadConfig() {
	if dockerEnv := os.Getenv("DOCKER_ENV"); dockerEnv == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env: ", err)
		}
	}

	BasePath = os.Getenv("BASE_PATH")
	Port = os.Getenv("PORT")

	Algorithm = os.Getenv("ALGORITHM")
	JwtExpireMins, _ = strconv.Atoi(os.Getenv("JWT_EXPIRE_MINUTES"))
	Secret = []byte(os.Getenv("SECRET"))
	WebApp = os.Getenv("WEBAPP")
	ApiKey = os.Getenv("API_KEY")

	DbDriver = os.Getenv("DB_DRIVER")
	DbUrl = os.Getenv("DATABASE_URL")
}
