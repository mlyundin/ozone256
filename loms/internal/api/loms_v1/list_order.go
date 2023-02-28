package loms_v1

import (
	"context"

	desc "route256/loms/pkg/loms_v1"
)

func (impl *Implementation) ListOrder(ctx context.Context, req *desc.ListOrderRequest) (*desc.ListOrderResponse, error) {
	order, err := impl.service.ListOrder(ctx, req.GetOrderId())

	if err != nil {
		return nil, err
	}

	items := make([]*desc.Item, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, &desc.Item{Sku: item.Sku, Count: item.Sku})
	}

	return &desc.ListOrderResponse{Status: order.Status, User: order.User, Items: items}, nil
}
