package main

import (
	"log"
	"net/http"
	"os"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/internal/web"
	"github.com/user-xat/short-link/pkg/middleware"
	"github.com/user-xat/short-link/pkg/templates"
)

type AppDeps struct {
	*configs.WebConfig
	errorLog *log.Logger
	infoLog  *log.Logger
}

func App(deps AppDeps) http.Handler {
	tmplCache, err := templates.NewTemplateCache(deps.WebConfig.HtmlTemplDir)
	if err != nil {
		log.Fatalf("can't create template cache using dir %s: %v", deps.WebConfig.HtmlTemplDir, err)
	}

	router := http.NewServeMux()

	webService := web.NewWebService(web.WebServiceDeps{
		WebConfig: deps.WebConfig,
		ErrorLog:  deps.errorLog,
		InfoLog:   deps.infoLog,
	})

	web.NewWebHandler(router, web.WebHandlerDeps{
		WebService:    webService,
		WebConfig:     deps.WebConfig,
		TemplateCache: tmplCache,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.Logging,
	)
	return stack(router)
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	conf := configs.LoadWebConfig()
	app := App(AppDeps{
		errorLog:  errorLog,
		infoLog:   infoLog,
		WebConfig: conf,
	})

	srv := http.Server{
		Addr:     ":" + conf.LaunchPort,
		ErrorLog: errorLog,
		Handler:  app,
	}

	infoLog.Printf("Server launch on http://localhost:%v", conf.LaunchPort)
	errorLog.Fatal(srv.ListenAndServe())
}
