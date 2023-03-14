-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS reservations(
    order_id bigint,
    warehouse_id bigint,
    sku bigint,
    count int,
    PRIMARY KEY (order_id, warehouse_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reservations;
-- +goose StatementEnd
