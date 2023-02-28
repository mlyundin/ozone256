package loms_v1

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	desc "route256/loms/pkg/loms_v1"
)

func (impl *Implementation) CancelOrder(ctx context.Context, req *desc.CancelOrderRequest) (*emptypb.Empty, error) {
	err := impl.service.CancelOrder(ctx, req.GetOrderId())
	return &emptypb.Empty{}, err
}
