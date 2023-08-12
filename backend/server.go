package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kotonohako/all-in/backend/generated/buf/kotobako/v1/kotobakov1connect"
	"github.com/kotonohako/all-in/backend/registry"
	"github.com/rs/cors"
)

func main() {
	api := http.NewServeMux()
	{
		kotobakoServer := &registry.KotobakoRegistry{}
		path, service := kotobakov1connect.NewKotobakoServiceHandler(kotobakoServer)
		api.Handle(path, service)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", api))
	addr := os.Getenv("BIND_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("Listening on %s \n", addr)
	http.ListenAndServe(addr,
		cors.AllowAll().Handler(mux),
	)
}
