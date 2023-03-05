package loms_client

import (
	"context"
	productServiceAPI "route256/product/pkg/product"

	"google.golang.org/grpc"
)

var _ Client = (*client)(nil)

type Product struct {
	Name  string
	Price uint32
}

type Client interface {
	GetProduct(ctx context.Context, token string, sku uint32) (*Product, error)
	ListSkus(ctx context.Context, token string, startAfterSkuu uint32, count uint32) ([]uint32, error)
}

type client struct {
	noteClient productServiceAPI.ProductServiceClient
}

func New(cc *grpc.ClientConn) *client {
	return &client{
		noteClient: productServiceAPI.NewProductServiceClient(cc),
	}
}

func (c *client) GetProduct(ctx context.Context, token string, sku uint32) (*Product, error) {
	res, err := c.noteClient.GetProduct(ctx, &productServiceAPI.GetProductRequest{Token: token, Sku: sku})
	if err != nil {
		return nil, err
	}

	return &Product{Name: res.Name, Price: res.Price}, nil
}

func (c *client) ListSkus(ctx context.Context, token string, startAfterSku uint32, count uint32) ([]uint32, error) {

	res, err := c.noteClient.ListSkus(ctx, &productServiceAPI.ListSkusRequest{Token: token, StartAfterSku: startAfterSku, Count: count})

	if err != nil {
		return nil, err
	}

	return res.GetSkus(), nil
}
