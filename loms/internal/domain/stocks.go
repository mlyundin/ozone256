package domain

import (
	"context"
	"route256/loms/pkg/model"
)

func (s *domainmodel) Stocks(ctx context.Context, sku uint32) ([]*model.StockItem, error) {
	return s.lomsRepo.Stocks(ctx, sku)
}
