package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/internal/auth"
	"github.com/user-xat/short-link/internal/cache"
	"github.com/user-xat/short-link/internal/link"
	"github.com/user-xat/short-link/internal/stat"
	"github.com/user-xat/short-link/internal/user"
	"github.com/user-xat/short-link/pkg/db"
	"github.com/user-xat/short-link/pkg/event"
	"github.com/user-xat/short-link/pkg/middleware"
	pb "github.com/user-xat/short-link/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	log.Println(conf.ServiceAddr)
	grpcClient, err := grpc.NewClient(conf.ServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed open grpc connection: %v", err)
	}

	// Loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	service := pb.NewShortLinkClient(grpcClient)
	database := db.NewDb(&conf.Db)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	cache, err := cache.NewCache(cache.CacheDeps{
		Config: cache.CacheConfig{
			Addr: conf.Cache.SocketAddress,
		},
		Ctx:      context.Background(),
		TTL:      5 * time.Minute,
		EventBus: eventBus,
	})
	if err != nil {
		log.Fatal(err)
	}

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
		Service:        service,
		ErrorLog:       errorLog,
		InfoLog:        infoLog,
		Cache:          cache,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         conf,
	})

	eventBus.Subscribe(statService.AddClick)
	eventBus.Subscribe(cache.UpdateCache)
	go eventBus.Start()

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}
