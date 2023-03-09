package domain

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

func (m *Model) AddToCart(ctx context.Context, item *CartItem) error {
	stocks, err := m.stocksChecker.Stocks(ctx, item.Sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks")
	}

	counter := int64(item.Count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			return nil
		}
	}

	return ErrInsufficientStocks
}
