package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/user-xat/short-link-server/pkg/models"
	"github.com/user-xat/short-link-server/pkg/models/memcached"
	pb "github.com/user-xat/short-link-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

	templCache, err := newTemplateCache(htmlTemplatesDir)
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

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.tmpl", nil)
}

func (app *application) createShortLinkHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !r.Form.Has("url") || r.Form.Get("url") == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	sourceLink := r.Form.Get("url")

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	link, err := app.serviceClient.Add(ctx, &wrapperspb.StringValue{Value: sourceLink})
	if err != nil {
		app.serverError(w, fmt.Errorf("remote service: %v", err))
		return
	}

	td := &templateData{Link: &models.LinkData{
		Source: link.Source,
		Short:  fmt.Sprintf("http://%s/%s", r.Host, link.Short),
	}}

	app.render(w, r, "home.page.tmpl", td)
}

func (app *application) shortLinkHandler(w http.ResponseWriter, r *http.Request) {
	shortlink := r.PathValue("shortlink")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if link, err := app.cache.Get(ctx, shortlink); err == nil {
		http.Redirect(w, r, link.Source, http.StatusSeeOther)
		return
	}

	link, err := app.serviceClient.Get(ctx, &wrapperspb.StringValue{Value: shortlink})
	if err != nil {
		app.serverError(w, err)
		return
	}

	_, err = app.cache.Set(context.Background(), &models.LinkData{
		Short:  link.Short,
		Source: link.Source,
	})
	if err != nil {
		app.errorLog.Printf("failed save value to cache: %v", err)
	}

	http.Redirect(w, r, link.Source, http.StatusSeeOther)
}
