package loms_v1

import (
	"context"
	desc "route256/loms/pkg/loms_v1"
)

func (impl *Implementation) Stocks(ctx context.Context, req *desc.StocksRequest) (*desc.StocksResponse, error) {
	stocks, err := impl.service.Stocks(ctx, req.GetSku())

	if err != nil {
		return nil, err
	}

	resStocks := make([]*desc.StockItem, 0, len(stocks))
	for _, stock := range stocks {
		resStocks = append(resStocks, &desc.StockItem{Count: stock.Count, WarehouseID: stock.WarehouseID})
	}

	return &desc.StocksResponse{Stocks: resStocks}, nil
}
