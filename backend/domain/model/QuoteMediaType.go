package model

type QuoteMediaType string

const (
	Tv           = QuoteMediaType("テレビ")
	Movie        = QuoteMediaType("映画")
	NobelOrManga = QuoteMediaType("小説・漫画")
	Commercial   = QuoteMediaType("広告")
	Others       = QuoteMediaType("その他")
)
