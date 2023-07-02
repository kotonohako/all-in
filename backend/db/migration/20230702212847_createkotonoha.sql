-- +goose Up
-- +goose StatementBegin
CREATE TABLE koto.kotonoha(
    kotonoha_id bigint,
    sentence text NOT NULL,
    author text NOT NULL,
    quote_source_name text,
    quote_media_type text,
    updated_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(kotonoha_id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE kotonoha;
-- +goose StatementEnd
