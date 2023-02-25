package handlers

import (
	"context"
	"log"
)

type CancelOrderRequest struct {
	OrderID int64 `json:"orderID"`
}

func (r CancelOrderRequest) Validate() error {
	return nil
}

type CancelOrderResponse struct {
}

type CancelOrderHandler struct{}

func (h *CancelOrderHandler) Handle(ctx context.Context, request CancelOrderRequest) (CancelOrderResponse, error) {
	log.Printf("cancelOrder: %+v", request)
	return CancelOrderResponse{}, nil
}
