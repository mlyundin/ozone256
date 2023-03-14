-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS carts (
    user_id bigint,
    sku bigint,
    count integer,
    constraint count check (count >= 0),
    PRIMARY KEY (user_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS carts;
-- +goose StatementEnd