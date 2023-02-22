package domain

import (
	"context"
)

func (m *Model) Purchase(ctx context.Context, userID int64) error {
	_, err := m.stocksChecker.CreateOrder(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
