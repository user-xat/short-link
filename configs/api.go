package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	Port        string
	Db          DbConfig
	Cache       DbConfig
	Auth        AuthConfig
	ServiceAddr string
}

type DbConfig struct {
	Dsn           string
	SocketAddress string
}

type AuthConfig struct {
	Secret string
}

func LoadApiConfig() *ApiConfig {
	var s *envStore
	values, err := godotenv.Read()
	if err != nil {
		log.Println("Error opening .env file. Default values are used")
		s = newEnvStore(nil)
	} else {
		s = newEnvStore(values)
	}
	return &ApiConfig{
		Port: s.getValue("API_PORT", "9090"),
		Db: DbConfig{
			Dsn: s.getValue("API_DSN", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"),
		},
		Auth: AuthConfig{
			Secret: s.getValue("API_SECRET", "my-256-bit-secret"),
		},
		Cache: DbConfig{
			SocketAddress: s.getValue("API_CACHE_ADDR", "redis:6379"),
		},
		ServiceAddr: s.getValue("API_SERVICE_ADDR", "service:9091"),
	}
}
