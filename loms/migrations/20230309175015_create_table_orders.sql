-- +goose Up
-- +goose StatementBegin
-- DROP TYPE IF EXISTS order_status;
-- CREATE TYPE order_status AS ENUM ('new', 'awaiting payment', 'failed', 'payed', 'cancelled');
CREATE TABLE IF NOT EXISTS orders(
    order_id bigserial,
    status int,
    user_id bigint,
    PRIMARY KEY (order_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- DROP TYPE IF EXISTS order_status;
-- +goose StatementEnd
