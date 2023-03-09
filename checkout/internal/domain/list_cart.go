package domain

import (
	"context"
	"log"
)

type ProductDesc struct {
	Name  string
	Price uint32
}

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

	var total uint32
	for i, item := range cart.Items {
		desc, err := m.productChecker.Product(ctx, item.Sku)
		if err != nil {
			log.Printf("Could not get description fot sku(%d)", item.Sku)
			continue
		}

		cart.Items[i].Name = desc.Name
		cart.Items[i].Price = desc.Price

		total += desc.Price * uint32(item.Count)
	}
	cart.Total = total

	return cart, nil
}
