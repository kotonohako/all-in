package registry

import (
	"context"
	"fmt"
	"strconv"

	connect_go "github.com/bufbuild/connect-go"
	kotobakov1 "github.com/kotonohako/all-in/backend/generated/buf/kotobako/v1"
	"github.com/kotonohako/all-in/backend/generated/buf/kotobako/v1/kotobakov1connect"
	"github.com/kotonohako/all-in/backend/repository"
)

var _ kotobakov1connect.KotobakoServiceHandler = &KotobakoRegistry{}

type KotobakoRegistry struct{}

func (s *KotobakoRegistry) Health(context.Context, *connect_go.Request[kotobakov1.HealthRequest]) (*connect_go.Response[kotobakov1.HealthResponse], error) {
	return connect_go.NewResponse(&kotobakov1.HealthResponse{
		Status: "Hello, Mr Kotobako",
	}), nil
}

func (s *KotobakoRegistry) ListQuotes(context.Context, *connect_go.Request[kotobakov1.ListQuotesRequest]) (*connect_go.Response[kotobakov1.ListQuotesResponse], error) {
	quotes, err := repository.GetQuotes()
	if err != nil {
		panic(fmt.Sprintf("err occurred: %v", err))
	}
	items := []*kotobakov1.Quote{}
	for _, quote := range quotes {
		q := &kotobakov1.Quote{
			QuoteId:         strconv.Itoa(quote.ID),
			AuthorName:      *quote.SpeakerName,
			QuoteMediaType:  string(quote.QuoteMediaType),
			QuoteSourceName: quote.QuoteSourceName,
			Sentence:        quote.Sentence,
		}
		items = append(items, q)
	}
	response := kotobakov1.ListQuotesResponse{
		Quotes: items,
	}
	return connect_go.NewResponse(&response), nil
}