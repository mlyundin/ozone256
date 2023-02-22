package domain

import (
	"context"
	"github.com/pkg/errors"
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

func (m *Model) ListCart(ctx context.Context, user int64) (Cart, error) {
	cartSkus := []uint32{1076963,
		1148162,
		1625903,
		2618151,
		2956315,
		2958025,
		3596599,
		3618852,
		4288068,
		4465995} // TODO temp

	var cart Cart

	for _, sku := range cartSkus {
		desk, err := m.productChecker.Product(ctx, sku)
		if err != nil {
			return cart, errors.Wrap(err, "Getting product desctription")
		}

		cart.Items = append(cart.Items, Product{Sku: sku, Count: 1, ProductDesc: desk})
		cart.Total += desk.Price
	}

	return cart, nil
}
