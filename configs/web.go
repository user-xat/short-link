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
		LaunchPort:   values["WEB_LAUNCH_PORT"],
		StaticDir:    values["WEB_STATIC_DIR"],
		HtmlTemplDir: values["WEB_HTML_TEMPL_DIR"],
		ApiAddr:      values["WEB_API_ADDR"],
		Cache: CacheDb{
			SocketAddress: values["WEB_CACHE_ADDR"],
		},
	}
}
