package domain

import (
	"context"
	"route256/loms/pkg/model"
)

func (s *domainmodel) ListOrder(ctx context.Context, orderId int64) (*model.Order, error) {

	order, err := s.lomsRepo.GetOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	items, err := s.lomsRepo.GetReservations(ctx, orderId)

	if err != nil {
		return nil, err
	}

	temp := make(map[uint32]uint16)
	for _, item := range items {
		temp[item.Sku] += item.Count
	}
	orderItems := make([]*model.Item, 0, len(temp))

	for sku, count := range temp {
		orderItems = append(orderItems, &model.Item{Sku: sku, Count: count})
	}

	return &model.Order{Status: order.Status, User: order.User, Items: orderItems}, nil
}
