-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sources
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    feed_url   VARCHAR(255) NOT NULL,
    priority   INT          NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS sources_id_uindex on sources (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sources;
DROP INDEX sources_id_uindex;
-- +goose StatementEnd
