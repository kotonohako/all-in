// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package generated

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /v1/health)
	HealthCheck(ctx echo.Context) error
	// セリフ一覧 API
	// (GET /v1/quotes)
	QuoteList(ctx echo.Context) error
	// セリフ登録 API
	// (POST /v1/quotes)
	RegisterQuote(ctx echo.Context) error
	// セリフ詳細 API
	// (GET /v1/quotes/{quoteId})
	QuoteDetail(ctx echo.Context, quoteId int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// HealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.HealthCheck(ctx)
	return err
}

// QuoteList converts echo context to params.
func (w *ServerInterfaceWrapper) QuoteList(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.QuoteList(ctx)
	return err
}

// RegisterQuote converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterQuote(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterQuote(ctx)
	return err
}

// QuoteDetail converts echo context to params.
func (w *ServerInterfaceWrapper) QuoteDetail(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "quoteId" -------------
	var quoteId int

	err = runtime.BindStyledParameterWithLocation("simple", false, "quoteId", runtime.ParamLocationPath, ctx.Param("quoteId"), &quoteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter quoteId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.QuoteDetail(ctx, quoteId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/v1/health", wrapper.HealthCheck)
	router.GET(baseURL+"/v1/quotes", wrapper.QuoteList)
	router.POST(baseURL+"/v1/quotes", wrapper.RegisterQuote)
	router.GET(baseURL+"/v1/quotes/:quoteId", wrapper.QuoteDetail)

}
