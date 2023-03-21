package domain

import "context"

type Stock struct {
	WarehouseID int64
	Count       uint64
}

type CartItem struct {
	User  int64
	Sku   uint32
	Count uint16
}

type StocksChecker interface {
	Stocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, userID int64, cart *Cart) (int64, error)
}

type ProductChecker interface {
	Product(ctx context.Context, sku uint32) (ProductDesc, error)
	Products(ctx context.Context, skus []uint32) ProductsDesc
	Skus(ctx context.Context, startAfterSku uint32, count uint32) ([]uint32, error)
}

type CartHandler interface {
	AddToCart(ctx context.Context, item *CartItem) error
	DeleteFromCart(ctx context.Context, item *CartItem) error
	GetItemCount(ctx context.Context, userId int64, sku uint32) (uint16, error)
	ListCart(ctx context.Context, userId int64) (*Cart, error)
}

type Model struct {
	stocksChecker  StocksChecker
	productChecker ProductChecker
	cartHandler    CartHandler
}

func New(stocksChecker StocksChecker, productChecker ProductChecker, cartHandler CartHandler) *Model {
	return &Model{
		stocksChecker:  stocksChecker,
		productChecker: productChecker,
		cartHandler:    cartHandler,
	}
}

func NewMock(deps ...interface{}) *Model {
	model := Model{}

	for _, v := range deps {
		switch s := v.(type) {
		case StocksChecker:
			model.stocksChecker = s
		case CartHandler:
			model.cartHandler = s
		case ProductChecker:
			model.productChecker = s
		}
	}

	return &model
}
