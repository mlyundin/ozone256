package handlers

import (
	"context"
	"log"
)

type StocksRequest struct {
	SKU uint32 `json:"sku"`
}

func (r StocksRequest) Validate() error {
	// TODO: implement
	return nil
}

type StockItem struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []StockItem `json:"stocks"`
}

func (h *Handler) HandleStocks(ctx context.Context, request StocksRequest) (StocksResponse, error) {
	log.Printf("stocks: %+v", request)
	return StocksResponse{
		Stocks: []StockItem{
			{
				WarehouseID: 123,
				Count:       5,
			},
		},
	}, nil
}
