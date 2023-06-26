package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPost, http.MethodDelete},
	}))
	e.GET("/hello", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "*")
		return c.JSON(http.StatusOK, "{ 'status': 'Hello, World!' }")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
