package loms

import (
	"context"
	"route256/checkout/internal/domain"
	loms "route256/loms/pkg/client/grpc/loms-service"
	"route256/loms/pkg/model"
)

type Client struct {
	grpcClient loms.Client
}

func New(grpcClient loms.Client) *Client {
	return &Client{grpcClient}
}

func (c *Client) Stocks(ctx context.Context, sku uint32) ([]domain.Stock, error) {
	reqStocks, err := c.grpcClient.Stocks(ctx, sku)
	if err != nil {
		return nil, err
	}

	stocks := make([]domain.Stock, 0, len(reqStocks))
	for _, stock := range reqStocks {
		stocks = append(stocks, domain.Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		})
	}

	return stocks, nil
}

func (c *Client) CreateOrder(ctx context.Context, userID int64, cart *domain.Cart) (int64, error) {
	items := make([]*model.Item, 0, len(cart.Items))
	for _, item := range cart.Items {
		items = append(items, &model.Item{Sku: item.Sku, Count: item.Count})
	}

	orderId, err := c.grpcClient.CreateOrder(ctx, userID, items)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}
