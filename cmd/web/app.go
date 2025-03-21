package main

import (
	"fmt"
	"html/template"
	"log"

	"github.com/user-xat/short-link/pkg/models/memcached"
	"github.com/user-xat/short-link/pkg/templates"
	pb "github.com/user-xat/short-link/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	cache         *memcached.CachedLinkModel
	clientConn    *grpc.ClientConn
	serviceClient pb.ShortLinkClient
	templateCache map[string]*template.Template
}

func NewApplication(errorLog, infoLog *log.Logger, htmlTemplatesDir, remoteService string, cacheServers []string) (*application, error) {
	grpcClient, err := grpc.NewClient(remoteService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed open grpc connection: %v", err)
	}
	slClient := pb.NewShortLinkClient(grpcClient)

	templCache, err := templates.NewTemplateCache(htmlTemplatesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create template cache: %v", err)
	}

	return &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		clientConn:    grpcClient,
		serviceClient: slClient,
		cache:         memcached.NewCachedLinkModel(cacheServers...),
		templateCache: templCache,
	}, nil
}

// Close all opened connections
func (app *application) Close() {
	app.clientConn.Close()
}

// func (app *application) shortLinkHandler(w http.ResponseWriter, r *http.Request) {
// 	shortlink := r.PathValue("shortlink")
// 	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
// 	defer cancel()
// 	if link, err := app.cache.Get(ctx, shortlink); err == nil {
// 		http.Redirect(w, r, link.Source, http.StatusSeeOther)
// 		return
// 	}
// 	link, err := app.serviceClient.Get(ctx, &wrapperspb.StringValue{Value: shortlink})
// 	if err != nil {
// 		if e, ok := status.FromError(err); ok {
// 			switch e.Code() {
// 			case codes.NotFound:
// 				res.ClientError(w, http.StatusNotFound)
// 			default:
// 				res.ServerError(w, app.errorLog, err)
// 			}
// 		} else {
// 			res.ServerError(w, app.errorLog, err)
// 		}
// 		return
// 	}
// 	_, err = app.cache.Set(context.Background(), &models.LinkData{
// 		Short:  link.Short,
// 		Source: link.Source,
// 	})
// 	if err != nil {
// 		app.errorLog.Printf("failed save value to cache: %v", err)
// 	}
// 	http.Redirect(w, r, link.Source, http.StatusSeeOther)
// }
