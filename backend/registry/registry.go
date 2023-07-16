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
	kotonohas, err := repository.GetQuotes()
	if err != nil {
		panic(fmt.Sprintf("err occurred: %v", err))
	}

	kotonohasResponse := []generated.KotonohaResponse{}
	for _, kotonoha := range kotonohas {
		kotonohaResponse := generated.KotonohaResponse{
			Author:          kotonoha.Author,
			KotonohaId:      kotonoha.ID,
			QuoteMediaType:  kotonoha.QuoteSourceName,
			QuoteSourceName: kotonoha.QuoteSourceName,
			Sentence:        kotonoha.Sentence,
		}
		kotonohasResponse = append(kotonohasResponse, kotonohaResponse)
	}
	return ctx.JSON(http.StatusOK, kotonohasResponse)
}
