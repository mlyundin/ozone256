-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders ADD COLUMN IF NOT EXISTS creation_time bigint;
CREATE UNIQUE INDEX creation_time_idx ON orders (creation_time);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX creation_time_idx;
ALTER TABLE orders DROP COLUMN IF EXISTS creation_time;
-- +goose StatementEnd
