-- name: ListQuotes :many
SELECT * FROM quote;

-- name: GetQuote :one
SELECT * FROM quote WHERE id = ?;