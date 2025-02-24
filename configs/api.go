package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *ApiConfig {
	values, err := godotenv.Read()
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	return &ApiConfig{
		Db: DbConfig{
			Dsn: values["API_DSN"],
		},
		Auth: AuthConfig{
			Secret: values["API_SECRET"],
		},
	}
}
