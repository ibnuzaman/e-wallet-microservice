package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

var Env = map[string]string{}

func SetupConfig() {

	var err error
	Env, err = godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string, val string) string {
	if Env[key] == "" {
		return val
	}
	return Env[key]
}
