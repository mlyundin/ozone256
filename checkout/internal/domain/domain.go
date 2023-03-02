package domain

import "context"

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, userID int64) (int64, error)
}

type ProductChecker interface {
	Product(ctx context.Context, sku uint32) (ProductDesc, error)
	Skus(ctx context.Context, startAfterSku uint32, count uint32) ([]uint32, error)
}

type Model struct {
	stocksChecker  StocksChecker
	productChecker ProductChecker
}

func New(stocksChecker StocksChecker, productChecker ProductChecker) *Model {
	return &Model{
		stocksChecker:  stocksChecker,
		productChecker: productChecker,
	}
}
