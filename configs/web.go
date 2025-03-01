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
}

func LoadWebConfig() *WebConfig {
	values, err := godotenv.Read()
	if err != nil {
		log.Println("Error opening .env file. Default values are used")
	}
	s := envStore(values)
	return &WebConfig{
		LaunchPort:   s.getValue("WEB_LAUNCH_PORT", "8110"),
		StaticDir:    s.getValue("WEB_STATIC_DIR", "./static"),
		HtmlTemplDir: s.getValue("WEB_HTML_TEMPL_DIR", "./html"),
		ApiAddr:      s.getValue("WEB_API_ADDR", "api:9090"),
	}
}
