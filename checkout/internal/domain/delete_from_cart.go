package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *Model) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint16) error {
	_, err := m.stocksChecker.Stocks(ctx, sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks")
	}

	return nil
}
