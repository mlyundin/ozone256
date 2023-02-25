package handlers

import (
	"context"
	"log"
)

type OrderPayedRequest struct {
	OrderID int64 `json:"orderID"`
}

func (r OrderPayedRequest) Validate() error {
	if r.OrderID == 0 {
		return ErrEmptyOrder
	}
	return nil
}

type OrderPayedResponse struct {
}

func (h *Handler) HandleOrderPayed(ctx context.Context, request OrderPayedRequest) (OrderPayedResponse, error) {
	log.Printf("orderPayed: %+v", request)
	return OrderPayedResponse{}, nil
}
