-- +migrate Up
CREATE TABLE IF NOT EXISTS "orders"
(
    id                   UUID             NOT NULL UNIQUE PRIMARY KEY,
    type                 VARCHAR(64)      NOT NULL,
    status               VARCHAR(32)      NOT NULL DEFAULT 'new',
    provider             VARCHAR(16)      NOT NULL,
    user_id              BIGINT           NOT NULL,
    wallet_address       VARCHAR(45)      NOT NULL,
    base_currency        VARCHAR(16)      NOT NULL DEFAULT 'EUR',
    base_currency_amount BIGINT           NOT NULL,
    base_currency_price  DOUBLE PRECISION NOT NULL,
    pay_currency         VARCHAR(16)      NOT NULL,
    pay_currency_amount  BIGINT           NOT NULL,
    pay_currency_price   DOUBLE PRECISION NOT NULL,
    rate_pay_to_base     DOUBLE PRECISION NOT NULL,
    coins                BIGINT           NOT NULL,
    fee                  BIGINT           NOT NULL,
    created_at           BIGINT           NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at           BIGINT           NOT NULL DEFAULT 0
);

CREATE INDEX orders_id_idx ON orders (id);

-- +migrate Down
DROP TABLE IF EXISTS "orders" CASCADE;
