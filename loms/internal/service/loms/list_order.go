package loms

import (
	"context"

	"route256/loms/pkg/model"
)

func (s *service) ListOrder(ctx context.Context, orderId int64) (*model.Order, error) {
	return &model.Order{Status: "created", User: 3, Items: []*model.Item{{Sku: 3, Count: 4}, {Sku: 5, Count: 100}}}, nil
}
