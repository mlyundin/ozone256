package loms_v1

import (
	"context"

	"route256/loms/internal/model"
	desc "route256/loms/pkg/loms_v1"
)

func (impl *Implementation) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	reqItems := req.GetItems()
	items := make([]*model.Item, 0, len(reqItems))

	for _, item := range reqItems {
		items = append(items, &model.Item{Sku: item.GetSku(), Count: uint16(item.GetCount())})
	}

	ordId, err := impl.service.CreateOrder(ctx, req.GetUser(), items)

	if err != nil {
		return nil, err
	}

	return &desc.CreateOrderResponse{OrderId: ordId}, nil
}
