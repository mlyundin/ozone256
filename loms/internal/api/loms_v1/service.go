package loms_v1

import (
	"context"
	"route256/loms/internal/domain"
	desc "route256/loms/pkg/loms_v1"
	"route256/loms/pkg/model"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type Implementation struct {
	desc.UnimplementedLomsV1Server

	domain domain.Model
}

func New(domain domain.Model) *Implementation {
	return &Implementation{
		desc.UnimplementedLomsV1Server{},
		domain,
	}
}

func (impl *Implementation) OrderPayed(ctx context.Context, req *desc.OrderPayedRequest) (*emptypb.Empty, error) {
	err := impl.domain.OrderPayed(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (impl *Implementation) ListOrder(ctx context.Context, req *desc.ListOrderRequest) (*desc.ListOrderResponse, error) {
	order, err := impl.domain.ListOrder(ctx, req.GetOrderId())

	if err != nil {
		return nil, err
	}

	items := make([]*desc.Item, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, &desc.Item{Sku: item.Sku, Count: item.Sku})
	}

	return &desc.ListOrderResponse{Status: order.Status, User: order.User, Items: items}, nil
}

func (impl *Implementation) CreateOrder(ctx context.Context, req *desc.CreateOrderRequest) (*desc.CreateOrderResponse, error) {
	reqItems := req.GetItems()
	items := make([]*model.Item, 0, len(reqItems))

	for _, item := range reqItems {
		items = append(items, &model.Item{Sku: item.GetSku(), Count: uint16(item.GetCount())})
	}

	ordId, err := impl.domain.CreateOrder(ctx, req.GetUser(), items)

	if err != nil {
		return nil, err
	}

	return &desc.CreateOrderResponse{OrderId: ordId}, nil
}

func (impl *Implementation) CancelOrder(ctx context.Context, req *desc.CancelOrderRequest) (*emptypb.Empty, error) {
	err := impl.domain.CancelOrder(ctx, req.GetOrderId())
	return &emptypb.Empty{}, err
}

func (impl *Implementation) Stocks(ctx context.Context, req *desc.StocksRequest) (*desc.StocksResponse, error) {
	stocks, err := impl.domain.Stocks(ctx, req.GetSku())

	if err != nil {
		return nil, err
	}

	resStocks := make([]*desc.StockItem, 0, len(stocks))
	for _, stock := range stocks {
		resStocks = append(resStocks, &desc.StockItem{Count: stock.Count, WarehouseID: stock.WarehouseID})
	}

	return &desc.StocksResponse{Stocks: resStocks}, nil
}
