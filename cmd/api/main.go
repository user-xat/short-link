package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/internal/auth"
	"github.com/user-xat/short-link/internal/link"
	"github.com/user-xat/short-link/internal/stat"
	"github.com/user-xat/short-link/internal/user"
	"github.com/user-xat/short-link/pkg/db"
	"github.com/user-xat/short-link/pkg/event"
	"github.com/user-xat/short-link/pkg/middleware"
)

func main() {
	conf := configs.LoadApiConfig()
	app := App(conf)
	server := http.Server{
		Addr:    ":" + conf.Port,
		Handler: app,
	}

	fmt.Printf("Server is listening on port %s\n", conf.Port)
	log.Fatal(server.ListenAndServe())
}

func App(conf *configs.ApiConfig) http.Handler {
	database := db.NewDb(&conf.Db)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(database)
	userRepository := user.NewUserRepository(database)
	statRepository := stat.NewStatRepository(database)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		ApiConfig:   conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	go statService.AddClick()

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}
