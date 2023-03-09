-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stock (
    warehouse_id bigint,
    sku bigint,
    count bigint,
    constraint count check (count >= 0),
    PRIMARY KEY (warehouse_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stock;
-- +goose StatementEnd