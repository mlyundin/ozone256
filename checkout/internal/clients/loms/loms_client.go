package loms

import (
	"context"
	"log"
	"route256/checkout/internal/domain"
	lomsClient "route256/loms/pkg/client/grpc/loms-service"
	"route256/loms/pkg/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcClient lomsClient.Client
}

func New(ctx context.Context, address string) *Client {

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	grpcClient := lomsClient.New(conn)
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

func (c *Client) CreateOrder(ctx context.Context, userID int64) (int64, error) { // TODO add params
	orderId, err := c.grpcClient.CreateOrder(ctx, userID, []*model.Item{})
	if err != nil {
		return 0, err
	}

	return orderId, nil
}
