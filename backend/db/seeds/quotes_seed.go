package main

import (
	"fmt"

	"github.com/kotonohako/all-in/backend/repository"
)

type dao struct {
	Sentence        string
	Author          string
	QuoteSourceName string `db:"quote_source_name"`
	QuoteMediaType  string `db:"quote_media_type"`
}

func main() {
	arr := []dao{
		{
			Sentence:        "わかんないものは わかんないまま取っておけばいい",
			Author:          "鈴木敏夫",
			QuoteSourceName: "鈴木敏夫自伝",
			QuoteMediaType:  "小説",
		},
		{
			Sentence:        "すべての女性には、嘘をつくための特別な独立器官のようなものが生まれつき具わっている",
			Author:          "村上春樹",
			QuoteSourceName: "独立器官",
			QuoteMediaType:  "小説",
		},
		{
			Sentence:        "職業とは本来、愛のある行為であるべきなんです。便宜的な結婚みたいなものじゃなく",
			Author:          "村上春樹",
			QuoteSourceName: "ノルウェイの森",
			QuoteMediaType:  "小説",
		},
	}

	query := `INSERT INTO quote(sentence, author, quote_source_name, quote_media_type) VALUES (:sentence, :author, :quote_source_name, :quote_media_type)`
	db, err := repository.DbConnection()
	if err != nil {
		panic(fmt.Sprintf("db connection can't get: %s", err))
	}

	_, err = db.NamedExec(query, arr)
	if err != nil {
		panic(fmt.Sprintf("can't bulk insert %s, caused by %s", query, err))
	}
}
