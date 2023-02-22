package products

import (
	"context"
	"route256/checkout/internal/domain"
	"route256/libs/httpclient"
)

type Client struct {
	url           string
	urlGetProduct string
	urlListSkus   string
	token         string
}

func New(url string, token string) *Client {
	return &Client{
		url:           url,
		urlGetProduct: url + "/get_product",
		urlListSkus:   url + "/get_skus",
		token:         token,
	}
}

type GetProductRequest struct {
	Token string `json:"token"`
	Sku   uint32 `json:"sku"`
}

type GetProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (c *Client) GetProduct(ctx context.Context, sku uint32) (domain.ProductDesc, error) {
	response, err := httpclient.Send[GetProductRequest, GetProductResponse](ctx, c.urlGetProduct,
		GetProductRequest{Token: c.token, Sku: sku})
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

func (c *Client) GetSkus(ctx context.Context, startAfterSku uint32, count uint32) error {
	_, err := httpclient.Send[GetSkusRequest, GetSkustResponse](ctx, c.urlListSkus,
		GetSkusRequest{Token: c.token, StartAfterSku: startAfterSku, Count: count}) // TODO handle response
	if err != nil {
		return err
	}

	return nil
}
