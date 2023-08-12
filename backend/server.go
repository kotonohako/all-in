package main

import (
	"context"
	"log"
	"net/http"
	"os"

	connect_go "github.com/bufbuild/connect-go"
	kotobakov1 "github.com/kotonohako/all-in/backend/generated/buf/kotobako/v1"
	"github.com/kotonohako/all-in/backend/generated/buf/kotobako/v1/kotobakov1connect"
	"github.com/rs/cors"
)

type KotobakoServer struct{}

func (s *KotobakoServer) Health(context.Context, *connect_go.Request[kotobakov1.HealthRequest]) (*connect_go.Response[kotobakov1.HealthResponse], error) {
	return connect_go.NewResponse(&kotobakov1.HealthResponse{
		Status: "Hello, Mr Kotobako",
	}), nil
}

func main() {
	api := http.NewServeMux()
	{
		kotobakoServer := &KotobakoServer{}
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
