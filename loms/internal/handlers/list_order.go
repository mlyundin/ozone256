package handlers

import (
	"context"
	"log"
)

type ListOrderRequest struct {
	OrderID int64 `json:"orderID"`
}

func (r ListOrderRequest) Validate() error {
	if r.OrderID == 0 {
		return ErrEmptyOrder
	}
	return nil
}

type ListOrderResponse struct {
	Status string `json:"status"`
	User   int64  `json:"user"`
	Items  []Item `json:"items"`
}

func (h *Handler) HandleListOrder(ctx context.Context, request ListOrderRequest) (ListOrderResponse, error) {
	log.Printf("listOrder: %+v", request)
	return ListOrderResponse{
		User:   12345,
		Status: "Open",
		Items: []Item{
			{
				Sku:   123,
				Count: 5,
			},
		},
	}, nil
}
