package domain

import (
	"context"
	"route256/libs/logger"

	"go.uber.org/zap"
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

	logger.Info("Create order", zap.Int64("orderId", orderID))

	return nil
}
