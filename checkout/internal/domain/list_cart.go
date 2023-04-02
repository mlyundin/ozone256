package domain

import (
	"context"
	"route256/libs/logger"

	"go.uber.org/zap"
)

type ProductDesc struct {
	Name  string
	Price uint32
	Err   error
}

type ProductsDesc = map[uint32]ProductDesc

type Product struct {
	Sku   uint32
	Count uint16
	ProductDesc
}

type Cart struct {
	Items []Product
	Total uint32
}

func (m *Model) ListCart(ctx context.Context, user int64) (*Cart, error) {
	cart, err := m.cartHandler.ListCart(ctx, user)
	if err != nil {
		return nil, err
	}

	skus := make([]uint32, 0, len(cart.Items))
	for _, item := range cart.Items {
		skus = append(skus, item.Sku)
	}

	productsDesc := m.productChecker.Products(ctx, skus)

	var total uint32
	for i, item := range cart.Items {
		desc, found := productsDesc[item.Sku]
		if !found || desc.Err != nil {
			logger.Error("Could not get description for ", zap.Uint32("sku", item.Sku))
			continue
		}

		cart.Items[i].Name = desc.Name
		cart.Items[i].Price = desc.Price

		total += desc.Price * uint32(item.Count)
	}
	cart.Total = total

	return cart, nil
}
