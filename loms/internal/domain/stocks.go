package domain

import (
	"context"
	"log"
	"route256/loms/pkg/model"
)

func (s *domainmodel) Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error) {
	log.Println("Stocks service")
	res := []*model.StockItem{{WarehouseID: 1, Count: 1}, {WarehouseID: 2, Count: 2}}
	return res, nil
}
