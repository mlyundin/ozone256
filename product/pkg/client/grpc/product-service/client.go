package loms_client

import (
	"context"
	"log"
	"route256/libs/workerpool"
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
	GetProducts(ctx context.Context, token string, skus []uint32) map[uint32]ProductRes
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
	log.Println("Send GetProduct for sku", sku)
	res, err := c.noteClient.GetProduct(ctx, &productServiceAPI.GetProductRequest{Token: token, Sku: sku})
	if err != nil {
		return nil, err
	}

	return &Product{Name: res.Name, Price: res.Price}, nil
}

type ProductRes struct {
	Product *Product
	Err     error
}

type productreswithsku struct {
	ProductRes
	sku uint32
}

func (c *client) GetProducts(ctx context.Context, token string, skus []uint32) map[uint32]ProductRes {
	n := len(skus)
	amountWorkers := n
	if amountWorkers > 5 {
		amountWorkers = 5
	}

	wp := workerpool.New[uint32, productreswithsku](ctx, amountWorkers)
	defer wp.Close()

	tasks := make([]workerpool.Task[uint32, productreswithsku], 0, n)
	for _, sku := range skus {
		tasks = append(tasks, workerpool.Task[uint32, productreswithsku]{InArgs: sku, Callback: func(sku uint32) productreswithsku {
			product, err := c.GetProduct(ctx, token, sku)
			return productreswithsku{ProductRes: ProductRes{Product: product, Err: err}, sku: sku}
		}})
	}
	wp.Submit(ctx, tasks)

	result := make(map[uint32]ProductRes, n)
	output := wp.Output()
	for i := 0; i < len(skus); i++ {
		res, ok := <-output
		if !ok {
			log.Println("Chanel is closed")
			break
		}

		result[res.sku] = res.ProductRes
	}

	return result
}

func (c *client) ListSkus(ctx context.Context, token string, startAfterSku uint32, count uint32) ([]uint32, error) {
	res, err := c.noteClient.ListSkus(ctx, &productServiceAPI.ListSkusRequest{Token: token, StartAfterSku: startAfterSku, Count: count})

	if err != nil {
		return nil, err
	}

	return res.GetSkus(), nil
}
