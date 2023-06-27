package main

import (
	ga "kotonohako-backend/presentation/generated/api"
	gt "kotonohako-backend/presentation/generated/types"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Registry struct{}

func (r Registry) Health(ctx echo.Context) error {
	status := "200 ok"
	helloResponse := gt.HelloResponse{Status: &status}
	return ctx.JSON(http.StatusOK, helloResponse)
}

func main() {
	e := echo.New()
	r := Registry{}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPost, http.MethodDelete},
	}))
	ga.RegisterHandlers(e, r)
	e.Logger.Fatal(e.Start(":8080"))
}
