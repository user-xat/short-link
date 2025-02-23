package main

import (
	"fmt"
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
	app := App()
	server := http.Server{
		Addr:    ":9090",
		Handler: app,
	}

	fmt.Println("Server is listening on port 9090")
	server.ListenAndServe()
}

func App() http.Handler {
	conf := configs.LoadConfig()
	database := db.NewDb(conf)
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
		Config:      conf,
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
