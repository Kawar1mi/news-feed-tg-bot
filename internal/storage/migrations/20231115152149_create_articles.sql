-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS articles (
    id           BIGSERIAL PRIMARY KEY,
    source_id    BIGINT       NOT NULL,
    title        VARCHAR(255) NOT NULL,
    link         TEXT         NOT NULL UNIQUE,
    summary      TEXT,
    published_at TIMESTAMP    NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    posted_at    TIMESTAMP,
    CONSTRAINT fk_articles_source_id
        FOREIGN KEY (source_id)
            REFERENCES sources (id)
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS articles_id_uindex on articles (id);
CREATE UNIQUE INDEX IF NOT EXISTS articles_source_id_uindex on articles (source_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS articles;
DROP INDEX articles_id_uindex;
DROP INDEX articles_source_id_uindex;
-- +goose StatementEnd
