package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	Port string
	Db   DbConfig
}

func LoadServiceConfig() *ServiceConfig {
	values, err := godotenv.Read()
	if err != nil {
		log.Panicln("Error loading .env file")
	}
	return &ServiceConfig{
		Port: values["SERVICE_PORT"],
		Db: DbConfig{
			Dsn: values["SERVICE_DSN"],
		},
	}
}
