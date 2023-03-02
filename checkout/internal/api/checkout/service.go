package checkout

import (
	"context"
	"route256/checkout/internal/domain"
	desc "route256/checkout/pkg/checkout"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var _ desc.CheckoutServer = (*implementation)(nil)

type implementation struct {
	desc.UnimplementedCheckoutServer

	domain *domain.Model
}

func New(domain *domain.Model) *implementation {
	return &implementation{
		desc.UnimplementedCheckoutServer{},
		domain,
	}
}

func (impl *implementation) AddToCart(ctx context.Context, req *desc.AddToCartRequest) (*emptypb.Empty, error) {
	err := impl.domain.AddToCart(ctx, req.GetUser(), req.GetCount(), uint16(req.GetCount()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (impl *implementation) DeleteFromCart(ctx context.Context, req *desc.DeleteFromCartRequest) (*emptypb.Empty, error) {
	err := impl.domain.DeleteFromCart(ctx, req.GetUser(), req.GetCount(), uint16(req.GetCount()))

	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (impl *implementation) ListCart(ctx context.Context, req *desc.ListCartRequest) (*desc.ListCartResponse, error) {

	cart, err := impl.domain.ListCart(ctx, req.GetUser())

	if err != nil {
		return nil, err
	}

	itemResp := make([]*desc.Item, 0, len(cart.Items))
	for _, item := range cart.Items {
		itemResp = append(itemResp, &desc.Item{Sku: item.Sku, Count: uint32(item.Count), Name: item.Name, Price: item.Price})
	}

	return &desc.ListCartResponse{Items: itemResp, TotalPrice: cart.Total}, nil
}

func (impl *implementation) Purchase(ctx context.Context, req *desc.PurchaseRequest) (*emptypb.Empty, error) {
	err := impl.domain.Purchase(ctx, req.GetUser())
	return &emptypb.Empty{}, err
}
