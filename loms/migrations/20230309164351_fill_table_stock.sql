-- +goose Up
-- +goose StatementBegin
INSERT INTO stock (warehouse_id, sku, count) VALUES (2, 1076963, 79) ON CONFLICT (warehouse_id, sku) DO NOTHING;
INSERT INTO stock (warehouse_id, sku, count) VALUES (4, 1076963, 99) ON CONFLICT (warehouse_id, sku) DO NOTHING;
INSERT INTO stock (warehouse_id, sku, count) VALUES (1, 1148162, 25) ON CONFLICT (warehouse_id, sku) DO NOTHING;
INSERT INTO stock (warehouse_id, sku, count) VALUES (8, 1148162, 24) ON CONFLICT (warehouse_id, sku) DO NOTHING;
INSERT INTO stock (warehouse_id, sku, count) VALUES (5, 6245113, 57) ON CONFLICT (warehouse_id, sku) DO NOTHING;
INSERT INTO stock (warehouse_id, sku, count) VALUES (7, 6245113, 40) ON CONFLICT (warehouse_id, sku) DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM stock
-- +goose StatementEnd
