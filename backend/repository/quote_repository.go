package repository

import (
	"fmt"

	"github.com/kotonohako/all-in/backend/db/dao"
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
			author, 
			quote_source_name, 
			quote_media_type
		)
			VALUES (:$1, :$2, :$3, :$4)`
	_, err = db.Exec(query, sentence, author, quote_source_name, quote_media_type)

	if err != nil {
		return err
	}

	return nil
}

func GetQuotes() ([]dao.QuoteDAO, error) {
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

	arr := []dao.QuoteDAO{}
	for rows.Next() {
		var dao dao.QuoteDAO
		err := rows.StructScan(&dao)
		fmt.Printf("sentence: %s", dao.Sentence)
		if err != nil {
			return nil, err
		}
		arr = append(arr, dao)
	}
	fmt.Printf("arr : %v", arr)
	return arr, nil
}

func GetQuote(quoteID int) (dao.QuoteDAO, error) {
	quote := dao.QuoteDAO{}
	db, err := DbConnection()
	if err != nil {
		return quote, err
	}

	query := `SELECT * FROM quote WHERE id=?`
	err = db.Get(&quote, query, quoteID)

	return quote, err
}
