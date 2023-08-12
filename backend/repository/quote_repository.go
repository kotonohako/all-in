package repository

import (
	"context"
	"fmt"

	"github.com/kotonohako/all-in/backend/domain/model"
	"github.com/kotonohako/all-in/backend/generated/sqlc"
)

func CreateQuote(
	sentence string,
	author string,
	quote_source_name string,
	quote_media_type string,
) error {
	db, err := DbConnectionWithSqlx()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO quote (
			sentence, 
			speaker_name, 
			quote_source_name, 
			quote_media_type
		)
			VALUES (?, ?, ?, ?)`
	_, err = db.Exec(query, sentence, author, quote_source_name, quote_media_type)

	if err != nil {
		return err
	}

	return nil
}

func GetQuotes() ([]model.Quote, error) {
	ctx := context.Background()

	db, err := DbConnection()
	if err != nil {
		return nil, err
	}
	queries := sqlc.New(db)

	rows, err := queries.ListQuotes(ctx)
	if err != nil {
		return nil, err
	}

	arr := []model.Quote{}
	for _, row := range rows {
		if err != nil {
			return nil, err
		}
		quote := model.Quote{
			ID:              int(row.ID),
			Sentence:        row.Sentence,
			SpeakerName:     &row.SpeakerName.String,
			QuoteSourceName: row.QuoteSourceName,
			QuoteMediaType:  model.QuoteMediaType(row.QuoteMediaType),
			UpdatedAt:       row.UpdatedAt,
		}
		arr = append(arr, quote)
	}
	fmt.Printf("arr : %v", arr)
	return arr, nil
}

func GetQuote(quoteID int) (model.Quote, error) {
	ctx := context.Background()
	db, err := DbConnection()
	if err != nil {
		return model.Quote{}, err
	}
	queries := sqlc.New(db)

	quote, err := queries.GetQuote(ctx, int64(quoteID))
	if err != nil {
		return model.Quote{}, err
	}

	return model.Quote{
		ID:              int(quote.ID),
		Sentence:        quote.Sentence,
		SpeakerName:     &quote.SpeakerName.String,
		QuoteSourceName: quote.QuoteSourceName,
		QuoteMediaType:  model.QuoteMediaType(quote.QuoteMediaType),
		UpdatedAt:       quote.UpdatedAt,
	}, err
}
