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
		log.Println("Error opening .env file. Default values are used")
	}
	s := envStore(values)
	return &ServiceConfig{
		Port: s.getValue("SERVICE_PORT", "9091"),
		Db: DbConfig{
			Dsn: s.getValue("SERVICE_DSN", "host=localhost user=postgres password=my_pass dbname=shortlink port=5432 sslmode=disable"),
		},
	}
}
