package domain

import (
	"context"
	"route256/loms/pkg/model"
)

func (s *domainmodel) ListOrder(ctx context.Context, orderId int64) (*model.Order, error) {
	return &model.Order{Status: 3, User: 3, Items: []*model.Item{{Sku: 3, Count: 4}, {Sku: 5, Count: 100}}}, nil
}
