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
	var s *envStore
	values, err := godotenv.Read()
	if err != nil {
		log.Println("Error opening .env file. Default values are used")
		s = newEnvStore(nil)
	} else {
		s = newEnvStore(values)
	}
	return &ServiceConfig{
		Port: s.getValue("SERVICE_PORT", "9091"),
		Db: DbConfig{
			Dsn: s.getValue("SERVICE_DSN", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"),
		},
	}
}
