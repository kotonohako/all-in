package repository

import (
	"fmt"

	"github.com/kotonohako/all-in/backend/db/dao"
	"github.com/kotonohako/all-in/backend/domain/model"
)

func CreateQuote(
	sentence string,
	author string,
	quote_source_name string,
	quote_media_type string,
) error {
	db, err := DbConnection()
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
	db, err := DbConnection()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM quote`
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arr := []model.Quote{}
	for rows.Next() {
		var dao dao.QuoteDAO
		err := rows.StructScan(&dao)
		fmt.Printf("sentence: %s", dao.Sentence)
		if err != nil {
			return nil, err
		}
		quote := model.Quote{
			ID:              dao.ID,
			Sentence:        dao.Sentence,
			SpeakerName:     dao.SpeakerName,
			QuoteSourceName: dao.QuoteSourceName,
			QuoteMediaType:  model.QuoteMediaType(dao.QuoteMediaType),
			UpdatedAt:       dao.UpdatedAt,
		}
		arr = append(arr, quote)
	}
	fmt.Printf("arr : %v", arr)
	return arr, nil
}

func GetQuote(quoteID int) (model.Quote, error) {
	quote := dao.QuoteDAO{}
	db, err := DbConnection()
	if err != nil {
		return model.Quote{}, err
	}

	query := `SELECT * FROM quote WHERE id=?`
	err = db.Get(&quote, query, quoteID)

	return model.Quote{
		ID:              quote.ID,
		Sentence:        quote.Sentence,
		SpeakerName:     quote.SpeakerName,
		QuoteSourceName: quote.QuoteSourceName,
		QuoteMediaType:  model.QuoteMediaType(quote.QuoteMediaType),
		UpdatedAt:       quote.UpdatedAt,
	}, err
}
