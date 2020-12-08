-- +migrate Up
CREATE TABLE IF NOT EXISTS "emission"
(
    num                  SERIAL           NOT NULL,
    id                   UUID             NOT NULL UNIQUE PRIMARY KEY,
    type                 VARCHAR(64)      NOT NULL,
    gold_unit            VARCHAR(12)      NOT NULL DEFAULT 'grams',
    gold_amount          BIGINT           NOT NULL,
    base_currency        VARCHAR(16)      NOT NULL DEFAULT 'EUR',
    base_currency_amount BIGINT           NOT NULL,
    rate_base_to_unit    DOUBLE PRECISION NOT NULL,
    reference            VARCHAR(255)     NOT NULL DEFAULT '',
    details              JSONB            NOT NULL DEFAULT '{}',
    created_at           BIGINT           NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at           BIGINT           NOT NULL DEFAULT 0
);

CREATE INDEX emission_id_idx ON emission (id);

-- +migrate Down
DROP TABLE IF EXISTS "emission" CASCADE;
