package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	Port  string
	Db    DbConfig
	Cache DbConfig
	Auth  AuthConfig
}

type DbConfig struct {
	Dsn           string
	SocketAddress string
}

type AuthConfig struct {
	Secret string
}

func LoadApiConfig() *ApiConfig {
	values, err := godotenv.Read()
	if err != nil {
		log.Println("Error opening .env file. Default values are used")
	}
	s := envStore(values)
	return &ApiConfig{
		Port: s.getValue("API_PORT", "9090"),
		Db: DbConfig{
			Dsn: s.getValue("API_DSN", "host=localhost user=postgres password=my_pass dbname=shortlink port=5432 sslmode=disable"),
		},
		Auth: AuthConfig{
			Secret: s.getValue("API_SECRET", "my-256-bit-secret"),
		},
		Cache: DbConfig{
			SocketAddress: s.getValue("API_CACHE_ADDR", "redis:6379"),
		},
	}
}
