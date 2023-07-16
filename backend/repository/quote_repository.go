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
		INSERT INTO kotonoha (
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

func GetQuotes() ([]dao.KotonohaDAO, error) {
	db, err := DbConnection()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM kotonoha`
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	arr := []dao.KotonohaDAO{}
	for rows.Next() {
		var dao dao.KotonohaDAO
		err := rows.StructScan(&dao)
		fmt.Printf("kotonoha_id: %s", dao.Sentence)
		if err != nil {
			panic(fmt.Sprintf("query scan error, cause by %s", err))
			return nil, err
		}
		arr = append(arr, dao)
	}
	fmt.Printf("arr : %v", arr)
	return arr, nil
}
