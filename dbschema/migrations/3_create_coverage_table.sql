-- +migrate Up
CREATE TABLE "coverage"
(
    id          SERIAL NOT NULL PRIMARY KEY,
    emission_id UUID   NOT NULL REFERENCES emission (id) ON DELETE CASCADE,
    order_id    UUID   NOT NULL REFERENCES orders (id) ON DELETE CASCADE,
    amount      BIGINT NOT NULL,
    created_at  BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW()),
    updated_at  BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX coverage_gold_emission_id_idx ON coverage (emission_id);
CREATE INDEX coverage_order_id_idx ON coverage (order_id);

-- +migrate Down
DROP TABLE IF EXISTS "coverage" CASCADE;
