package loms

import (
	"context"
	"route256/loms/internal/model"
)

var _ Service = (*service)(nil)

type Service interface {
	Stoks(ctx context.Context, sku uint32) ([]*model.StockItem, error)

	CreateOrder(ctx context.Context, user int64, items []*model.Item) (int64, error)

	ListOrder(ctx context.Context, orderId int64) (*model.Order, error)

	CancelOrder(ctx context.Context, orderId int64) error

	OrderPayed(ctx context.Context, orderId int64) error
}

type service struct {
}

func NewService() *service {
	return &service{}
}
