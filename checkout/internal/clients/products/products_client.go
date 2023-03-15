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

func (c *Client) Products(ctx context.Context, skus []uint32) domain.ProductsDesc {
	response := c.grpcClient.GetProducts(ctx, c.token, skus)
	res := make(domain.ProductsDesc, len(response))

	for k, v := range response {
		res[k] = domain.ProductDesc{Name: v.Product.Name, Price: v.Product.Price, Err: v.Err}
	}

	return res
}

func (c *Client) Skus(ctx context.Context, startAfterSku uint32, count uint32) ([]uint32, error) {
	return c.grpcClient.ListSkus(ctx, c.token, startAfterSku, count)
}
