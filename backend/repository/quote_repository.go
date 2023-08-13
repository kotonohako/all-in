package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kotonohako/all-in/backend/domain/model"
	"github.com/kotonohako/all-in/backend/generated/sqlc"
)

func CreateQuote(
	sentence string,
	SpeakerName string,
	quoteSourceName string,
	quoteMediaType string,
) (uint, error) {
	ctx := context.Background()

	db, err := DbConnection()
	if err != nil {
		return 0, err
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err

	}
	defer tx.Rollback()

	queries := sqlc.New(db)
	qtx := queries.WithTx(tx)

	if err = qtx.CreateQuote(ctx,
		sqlc.CreateQuoteParams{
			Sentence:        sentence,
			SpeakerName:     sql.NullString{String: SpeakerName, Valid: true},
			QuoteSourceName: quoteSourceName,
			QuoteMediaType:  quoteMediaType,
		},
	); err != nil {
		return 0, err
	}

	queryId, err := qtx.GetLastInsertId(ctx)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return uint(queryId), nil
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
