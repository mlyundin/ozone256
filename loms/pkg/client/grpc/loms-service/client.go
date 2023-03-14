package loms_client

import (
	"context"
	lomsServiceAPI "route256/loms/pkg/loms"
	"route256/loms/pkg/model"

	"google.golang.org/grpc"
)

var _ Client = (*client)(nil)

type Client interface {
	Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error)

	CreateOrder(ctx context.Context, user int64, items []*model.Item) (int64, error)

	ListOrder(ctx context.Context, orderId int64) (*model.Order, error)

	CancelOrder(ctx context.Context, orderId int64) error

	OrderPayed(ctx context.Context, orderId int64) error
}

type client struct {
	noteClient lomsServiceAPI.LomsClient
}

func New(cc *grpc.ClientConn) *client {
	return &client{
		noteClient: lomsServiceAPI.NewLomsClient(cc),
	}
}

func (c *client) Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error) {
	res, err := c.noteClient.Stocks(ctx, &lomsServiceAPI.StocksRequest{Sku: sku})
	if err != nil {
		return nil, err
	}

	stocks := make([]*model.StockItem, 0, len(res.Stocks))

	for _, stock := range res.GetStocks() {
		stocks = append(stocks, &model.StockItem{WarehouseID: stock.GetWarehouseID(), Count: stock.GetCount()})
	}

	return stocks, nil
}

func (c *client) CreateOrder(ctx context.Context, user int64, items []*model.Item) (int64, error) {
	itemsReq := make([]*lomsServiceAPI.Item, 0, len(items))
	for _, item := range items {
		itemsReq = append(itemsReq, &lomsServiceAPI.Item{Sku: item.Sku, Count: uint32(item.Count)})
	}

	res, err := c.noteClient.CreateOrder(ctx, &lomsServiceAPI.CreateOrderRequest{User: user, Items: itemsReq})
	if err != nil {
		return 0, err
	}

	return res.GetOrderId(), nil
}

func (c *client) ListOrder(ctx context.Context, orderId int64) (*model.Order, error) {
	res, err := c.noteClient.ListOrder(ctx, &lomsServiceAPI.ListOrderRequest{OrderId: orderId})
	if err != nil {
		return nil, err
	}

	resItems := res.GetItems()
	items := make([]*model.Item, 0, len(resItems))
	for _, item := range resItems {
		items = append(items, &model.Item{Sku: item.GetSku(), Count: uint16(item.GetCount())})
	}
	return &model.Order{Status: model.OrderStatus(res.GetStatus()), User: res.GetUser(), Items: items}, nil
}

func (c *client) CancelOrder(ctx context.Context, orderId int64) error {
	_, err := c.noteClient.CancelOrder(ctx, &lomsServiceAPI.CancelOrderRequest{OrderId: orderId})
	return err
}

func (c *client) OrderPayed(ctx context.Context, orderId int64) error {
	_, err := c.noteClient.OrderPayed(ctx, &lomsServiceAPI.OrderPayedRequest{OrderId: orderId})
	return err
}
