package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"time"

	"github.com/user-xat/short-link-server/pkg/models/redis"
	pb "github.com/user-xat/short-link-server/proto"
	"google.golang.org/grpc"
)

var (
	port   = flag.String("port", "54321", "the service port")
	dbAddr = flag.String("db", "redis:6379", "the db address")
)

func main() {
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cfg := redis.Config{
		Addr:        *dbAddr,
		Password:    "",
		User:        "",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}

	store, err := redis.NewLinkStoreRedis(context.Background(), cfg)
	if err != nil {
		errorLog.Fatalf("failed to create redis store: %v", err)
	}

	addr := ":" + *port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errorLog.Fatalf("failed to listen: %s", addr)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterShortLinkServer(s, NewShortLinkService(NewShortLink(store)))

	infoLog.Printf("Starting gRPC listener at %v ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		errorLog.Fatalf("failed to serve: %v", err)
	}
}
