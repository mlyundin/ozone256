package domain

import (
	"context"
	"route256/loms/pkg/model"
)

var _ Model = (*domainmodel)(nil)

type Reservation struct {
	OrderId     int64
	WarehouseID int64
	Sku         uint32
	Count       uint16
}

type TransactionManager interface {
	RunRepeteableRead(ctx context.Context, f func(ctxTX context.Context) error) error
}

type LomsRepository interface {
	Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error)
	Reserve(ctx context.Context, reservation *Reservation) error
	Release(ctx context.Context, reservation *Reservation) error

	NewOrder(ctx context.Context, user int64) (int64, error)
	UpdateStatus(ctx context.Context, orderId int64, newStatus model.OrderStatus, currStatus model.OrderStatus) error
	GetOrder(ctx context.Context, orderId int64) (*model.Order, error)

	NewReservation(ctx context.Context, reservation *Reservation) error
	ReleaseAllReservations(ctx context.Context, orderId int64) error
	GetReservations(ctx context.Context, orderId int64) ([]*Reservation, error)
}

type Model interface {
	Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error)

	CreateOrder(ctx context.Context, user int64, items []*model.Item) (int64, error)

	ListOrder(ctx context.Context, orderId int64) (*model.Order, error)

	CancelOrder(ctx context.Context, orderId int64) error

	OrderPayed(ctx context.Context, orderId int64) error
}

type domainmodel struct {
	lomsRepo LomsRepository
	tm       TransactionManager
}

func New(lomsRepo LomsRepository, tm TransactionManager) *domainmodel {
	return &domainmodel{lomsRepo, tm}
}
