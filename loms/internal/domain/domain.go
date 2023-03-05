package domain

import (
	"context"
	"route256/loms/pkg/model"
)

var _ Model = (*domainmodel)(nil)

type Model interface {
	Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error)

	CreateOrder(ctx context.Context, user int64, items []*model.Item) (int64, error)

	ListOrder(ctx context.Context, orderId int64) (*model.Order, error)

	CancelOrder(ctx context.Context, orderId int64) error

	OrderPayed(ctx context.Context, orderId int64) error
}

type domainmodel struct {
}

func New() *domainmodel {
	return &domainmodel{}
}
