package registry

import (
	"fmt"
	"net/http"

	"github.com/kotonohako/all-in/backend/presentation/generated"
	"github.com/kotonohako/all-in/backend/repository"
	"github.com/labstack/echo/v4"
)

type ApiRegistry struct{}

func (r ApiRegistry) Health(ctx echo.Context) error {
	status := "200 ok"
	helloResponse := generated.HelloResponse{Status: &status}
	return ctx.JSON(http.StatusOK, helloResponse)
}

func (r ApiRegistry) API(ctx echo.Context) error {
	quotes, err := repository.GetQuotes()
	if err != nil {
		panic(fmt.Sprintf("err occurred: %v", err))
	}

	quotesResponse := []generated.KotonohaResponse{}
	for _, quote := range quotes {
		quoteResponse := generated.KotonohaResponse{
			Author:          quote.Author,
			KotonohaId:      quote.ID,
			QuoteMediaType:  quote.QuoteSourceName,
			QuoteSourceName: quote.QuoteSourceName,
			Sentence:        quote.Sentence,
		}
		quotesResponse = append(quotesResponse, quoteResponse)
	}
	return ctx.JSON(http.StatusOK, quotesResponse)
}
