package domain

import (
	"context"
)

type ProductDesc struct {
	Name  string
	Price uint32
}

type Product struct {
	Sku   uint32
	Count uint16
	ProductDesc
}

func (m *Model) ListCart(ctx context.Context, user int64) error {

	return ErrInsufficientStocks
}
