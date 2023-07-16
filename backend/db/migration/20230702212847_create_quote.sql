-- +goose Up
-- +goose StatementBegin
CREATE TABLE quote(
    id bigint NOT NULL AUTO_INCREMENT,
    sentence text NOT NULL,
    speaker_name text,
    quote_source_name text NOT NULL,
    quote_media_type text NOT NULL,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE quote;
-- +goose StatementEnd
