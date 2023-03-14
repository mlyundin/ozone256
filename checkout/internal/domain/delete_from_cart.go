package domain

import (
	"context"
)

func (m *Model) DeleteFromCart(ctx context.Context, item *CartItem) error {
	return m.cartHandler.DeleteFromCart(ctx, item)
}
