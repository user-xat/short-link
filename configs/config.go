package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	values, err := godotenv.Read()
	if err != nil {
		log.Panicln("Error loading .env file")
	}

	return &Config{
		Db: DbConfig{
			Dsn: values["DSN"],
		},
		Auth: AuthConfig{
			Secret: values["SECRET"],
		},
	}
}
