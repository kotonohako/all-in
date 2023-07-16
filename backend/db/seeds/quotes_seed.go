package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	"github.com/kotonohako/all-in/backend/repository"
)

type dao struct {
	Sentence        string `csv:"sentence"`
	SpeakerName     string `db:"speaker_name" csv:"speaker_name"`
	QuoteSourceName string `db:"quote_source_name" csv:"quote_source_name"`
	QuoteMediaType  string `db:"quote_media_type" csv:"quote_media_type"`
}

func main() {
	f, err := os.OpenFile(filepath.Join("db", "seeds", "csv", "quote_dummy.csv"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	arr := []dao{}

	if err := gocsv.UnmarshalFile(f, &arr); err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO quote(sentence, speaker_name, quote_source_name, quote_media_type) VALUES (:sentence, :speaker_name, :quote_source_name, :quote_media_type)`
	db, err := repository.DbConnection()
	if err != nil {
		panic(fmt.Sprintf("db connection can't get: %s", err))
	}

	_, err = db.NamedExec(query, arr)
	if err != nil {
		panic(fmt.Sprintf("can't bulk insert %s, caused by %s", query, err))
	}
}
