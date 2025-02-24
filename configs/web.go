package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type WebConfig struct {
	LaunchPort   string
	StaticDir    string
	HtmlTemplDir string
	ApiAddr      string
	Cache        CacheDb
}

type CacheDb struct {
	SocketAddress string
}

func LoadWebConfig() *WebConfig {
	values, err := godotenv.Read()
	if err != nil {
		log.Panicln("error loading .env file")
	}

	return &WebConfig{
		LaunchPort:   values["LAUNCH_PORT"],
		StaticDir:    values["STATIC_DIR"],
		HtmlTemplDir: values["HTML_TEMPL_DIR"],
		ApiAddr:      values["API_ADDR"],
		Cache: CacheDb{
			SocketAddress: values["CACHE_SOCKET_ADDR"],
		},
	}
}
