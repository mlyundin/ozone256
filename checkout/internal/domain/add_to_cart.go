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

	alreadyOrdered, err := m.cartHandler.GetItemCount(ctx, item.User, item.Sku)
	if err != nil {
		return errors.WithMessage(err, "checking already ordered")
	}

	counter := int64(item.Count) + int64(alreadyOrdered)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			err = m.cartHandler.AddToCart(ctx, item)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return ErrInsufficientStocks
}
