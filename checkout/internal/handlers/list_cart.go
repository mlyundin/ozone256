package handlers

import (
	"context"
	"log"
)

type ListCartRequest struct {
	User int64 `json:"user"`
}

func (r ListCartRequest) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}
	return nil
}

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type ListCartResponse struct {
	Items      []Item `json:"items"`
	TotalPrice uint32 `json:"totalPrice"`
}

func (h *Handler) HandleListCart(ctx context.Context, req ListCartRequest) (ListCartResponse, error) {
	log.Printf("listCart: %+v", req)

	var response ListCartResponse

	cart, err := h.businessLogic.ListCart(ctx, req.User)
	if err != nil {
		return response, err
	}

	response.Items = make([]Item, 0, len(cart.Items))
	for _, item := range cart.Items {
		response.Items = append(response.Items, Item{Sku: item.Sku, Name: item.Name, Price: item.Price, Count: item.Count})
	}
	response.TotalPrice = cart.Total

	return response, nil
}
