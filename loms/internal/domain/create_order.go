package domain

import (
	"context"
	"route256/loms/pkg/model"
)

func (s *domainmodel) CreateOrder(ctx context.Context, user int64, items []*model.Item) (int64, error) {
	return 1, nil
}
