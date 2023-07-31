package registry

import (
	"fmt"
	"net/http"

	"github.com/kotonohako/all-in/backend/presentation/generated"
	"github.com/kotonohako/all-in/backend/repository"
	"github.com/labstack/echo/v4"
)

type ApiRegistry struct{}

func (r ApiRegistry) HealthCheck(ctx echo.Context) error {
	status := "200 ok"
	helloResponse := generated.HelloResponse{Status: &status}
	return ctx.JSON(http.StatusOK, helloResponse)
}

func (r ApiRegistry) QuoteList(ctx echo.Context) error {
	quotes, err := repository.GetQuotes()
	if err != nil {
		panic(fmt.Sprintf("err occurred: %v", err))
	}

	quotesResponse := []generated.QuoteResponse{}
	for _, quote := range quotes {
		quoteResponse := generated.QuoteResponse{
			Id:              quote.ID,
			SpeakerName:     quote.SpeakerName,
			QuoteMediaType:  string(quote.QuoteMediaType),
			QuoteSourceName: quote.QuoteSourceName,
			Sentence:        quote.Sentence,
		}
		quotesResponse = append(quotesResponse, quoteResponse)
	}
	return ctx.JSON(http.StatusOK, quotesResponse)
}

func (r ApiRegistry) RegisterQuote(ctx echo.Context) error {
	q := new(generated.RegisterQuoteJSONRequestBody)
	if err := ctx.Bind(q); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	if err := repository.CreateQuote(q.Sentence, q.SpeakerName, q.QuoteSourceName, q.QuoteMediaType); err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("register by repository had err: %v", err))
	}

	return ctx.NoContent(204)
}

func (r ApiRegistry) QuoteDetail(ctx echo.Context, quoteId int) error {
	quote, err := repository.GetQuote(quoteId)
	if err != nil {
		panic(fmt.Sprintf("quote detail get error: %v", err))
	}

	quoteResponse := generated.QuoteResponse{
		Id:              quote.ID,
		SpeakerName:     quote.SpeakerName,
		QuoteMediaType:  string(quote.QuoteMediaType),
		QuoteSourceName: quote.QuoteSourceName,
		Sentence:        quote.Sentence,
	}

	return ctx.JSON(http.StatusOK, quoteResponse)
}
