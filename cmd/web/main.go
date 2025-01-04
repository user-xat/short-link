package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", "localhost", "addr for connection")
	port = flag.String("port", "8080", "port for server")
)

func main() {
	flag.Parse()
	s := NewServer(NewLinksStoreMap())
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", s.homeHandler)

	log.Printf("Server launch on http://%v:%v", *addr, *port)
	log.Fatal(http.ListenAndServe(*addr+":"+*port, mux))
}
