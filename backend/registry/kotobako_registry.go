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

func (s *KotobakoRegistry) GetQuote(context context.Context, request *connect_go.Request[kotobakov1.GetQuoteRequest]) (*connect_go.Response[kotobakov1.GetQuoteResponse], error) {
	quoteId, err := strconv.Atoi(request.Msg.QuoteId)
	if err != nil {
		return nil, fmt.Errorf("failed to convert requested quoteID %v", err)
	}

	quote, err := repository.GetQuote(quoteId)
	if err != nil {
		return nil, fmt.Errorf("quote detail get error: %v", err)
	}

	q := kotobakov1.GetQuoteResponse{
		QuoteId:         strconv.Itoa(quote.ID),
		AuthorName:      *quote.SpeakerName,
		QuoteMediaType:  string(quote.QuoteMediaType),
		QuoteSourceName: quote.QuoteSourceName,
		Sentence:        quote.Sentence,
	}

	return connect_go.NewResponse(&q), nil
}

func (s *KotobakoRegistry) PostQuote(context context.Context, request *connect_go.Request[kotobakov1.PostQuoteRequest]) (*connect_go.Response[kotobakov1.PostQuoteResponse], error) {
	msg := request.Msg
	sentence := msg.Sentence
	authorName := msg.AuthorName
	quoteSourceName := msg.QuoteSourceName
	quoteMediaType := msg.QuoteMediaType

	if err := repository.CreateQuote(sentence, authorName, quoteSourceName, quoteMediaType); err != nil {
		return nil, fmt.Errorf("failed to register quote, %v", err)
	}
	quoteId := "100"
	response := kotobakov1.PostQuoteResponse{
		QuoteId: quoteId,
	}
	return connect_go.NewResponse(&response), nil
}
