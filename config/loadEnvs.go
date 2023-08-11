package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvs() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}
}
