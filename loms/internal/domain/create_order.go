package domain

import (
	"context"
	"route256/loms/pkg/model"
)

func (s *domainmodel) CreateOrder(ctx context.Context, user int64, items []*model.Item) (int64, error) {
	orderId, err := s.lomsRepo.NewOrder(ctx, user)
	if err != nil {
		return orderId, err
	}

	return orderId, nil
}
