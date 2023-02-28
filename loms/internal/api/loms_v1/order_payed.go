package loms_v1

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
	desc "route256/loms/pkg/loms_v1"
)

func (impl *Implementation) OrderPayed(ctx context.Context, req *desc.OrderPayedRequest) (*emptypb.Empty, error) {
	err := impl.service.OrderPayed(ctx, req.GetOrderId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
