package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *Model) DeleteFromCart(ctx context.Context, item *CartItem) error {
	_, err := m.stocksChecker.Stocks(ctx, item.Sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks")
	}

	return nil
}
