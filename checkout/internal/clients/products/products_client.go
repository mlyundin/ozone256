package products

import (
	"context"
	"route256/checkout/internal/domain"
	product "route256/product/pkg/client/grpc/product-service"
)

type Client struct {
	grpcClient product.Client
	token      string
}

func New(grpcClient product.Client, token string) *Client {
	return &Client{
		grpcClient,
		token,
	}
}

func (c *Client) Product(ctx context.Context, sku uint32) (domain.ProductDesc, error) {
	response, err := c.grpcClient.GetProduct(ctx, c.token, sku)
	if err != nil {
		return domain.ProductDesc{}, err
	}

	return domain.ProductDesc{Name: response.Name, Price: response.Price}, nil
}

type GetSkusRequest struct {
	Token         string `json:"token"`
	StartAfterSku uint32 `json:"startAfterSku"`
	Count         uint32 `json:"count"`
}

type GetSkustResponse struct {
	Skus []uint32 `json:"Skus"`
}

func (c *Client) Skus(ctx context.Context, startAfterSku uint32, count uint32) ([]uint32, error) {
	return c.grpcClient.ListSkus(ctx, c.token, startAfterSku, count)
}
