package model

import "time"

type Quote struct {
	ID              int
	Sentence        string
	SpeakerName     *string
	QuoteSourceName string
	QuoteMediaType  QuoteMediaType
	UpdatedAt       time.Time
}
