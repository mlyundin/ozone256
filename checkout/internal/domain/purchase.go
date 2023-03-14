package domain

import (
	"context"
	"log"
)

func (m *Model) Purchase(ctx context.Context, userID int64) error {
	cart, err := m.cartHandler.ListCart(ctx, userID)
	if err != nil {
		return err
	}

	orderID, err := m.stocksChecker.CreateOrder(ctx, userID, cart)
	if err != nil {
		return err
	}

	log.Printf("Create order %d\n", orderID)

	return nil
}
