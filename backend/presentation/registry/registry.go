package registry

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiRegistry struct{}

func (r ApiRegistry) Health(ctx echo.Context) error {
	status := "200 ok"
	helloResponse := HelloResponse{Status: &status}
	return ctx.JSON(http.StatusOK, helloResponse)
}

func (r ApiRegistry) API(ctx echo.Context) error {
	quoteMediaType := "本"
	quoteSourceName := "ノルウェイの森"
	kotonohasResponse := KotonohaResponse{
		Author:          "村上春樹",
		KotonohaId:      12345,
		QuoteMediaType:  &quoteMediaType,
		QuoteSourceName: &quoteSourceName,
		Sentence:        "希望があるところには必ず試練があるものだから",
	}
	return ctx.JSON(http.StatusOK, kotonohasResponse)
}
