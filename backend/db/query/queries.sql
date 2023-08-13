-- name: ListQuotes :many
SELECT * FROM quote;

-- name: GetQuote :one
SELECT * FROM quote WHERE id = ?;

-- name: CreateQuote :exec
INSERT INTO quote (
    sentence, 
    speaker_name, 
    quote_source_name, 
    quote_media_type
)VALUES (?, ?, ?, ?);

-- name: GetLastInsertId :one
SELECT LAST_INSERT_ID();