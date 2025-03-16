package main

import (
	"log"
	"net"
	"os"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/internal/link"
	"github.com/user-xat/short-link/pkg/db"
	pb "github.com/user-xat/short-link/proto"
	"google.golang.org/grpc"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	conf := configs.LoadServiceConfig()

	// cfg := redis.Config{
	// 	Addr:        conf.Db.Dsn,
	// 	Password:    "",
	// 	User:        "",
	// 	DB:          0,
	// 	MaxRetries:  5,
	// 	DialTimeout: 10 * time.Second,
	// 	Timeout:     5 * time.Second,
	// }

	// store, err := redis.NewLinkStoreRedis(context.Background(), cfg)
	// if err != nil {
	// 	errorLog.Fatalf("failed to create redis store: %v", err)
	// }

	addr := ":" + conf.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errorLog.Fatalf("failed to listen: %s", addr)
	}
	defer lis.Close()

	database := db.NewDb(&conf.Db)
	linkRepository := link.NewLinkRepository(database)
	grpcHandler := link.NewGRPCHandler(link.GRPCHandlerDeps{
		LinkRepository: linkRepository,
	})

	s := grpc.NewServer()
	pb.RegisterShortLinkServer(s, grpcHandler)

	infoLog.Printf("Starting gRPC listener at %v ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		errorLog.Fatalf("failed to serve: %v", err)
	}
}
