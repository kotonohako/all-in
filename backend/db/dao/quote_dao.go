package dao

import "time"

type QuoteDAO struct {
	ID              int       `db:"id"`
	Sentence        string    `db:"sentence"`
	Author          string    `db:"author"`
	QuoteSourceName *string   `db:"quote_source_name"`
	QuoteMediaType  *string   `db:"quote_media_type"`
	UpdatedAt       time.Time `db:"updated_at"`
}