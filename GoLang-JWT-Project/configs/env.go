package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvVariables struct {
	DBConstraints      string `default:""`
	AuthServiceBaseUrl string `default:""`
}

var env *EnvVariables

func initEnvVariable() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	env = &EnvVariables{
		DBConstraints:      os.Getenv("DB_CONSTRAINTS"),
		AuthServiceBaseUrl: os.Getenv("AUTH_SERVICE_BASE_URL"),
	}
}

func GetEnv() *EnvVariables {
	initEnvVariable()
	if env == nil {
		log.Fatal("Error occurred during initEnvVariable method.")
	}
	return env
}
