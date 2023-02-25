package handlers

import (
	"context"
	"errors"
	"log"
)

type CreateOrderRequest struct {
	User  int64  `json:"user"`
	Items []Item `json:"items"`
}

var (
	ErrEmptyUser = errors.New("empty user")
)

func (r CreateOrderRequest) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}
	return nil
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

type CreateOrderHandler struct{}

func (h *CreateOrderHandler) Handle(ctx context.Context, request CreateOrderRequest) (CreateOrderResponse, error) {
	log.Printf("createOrder: %+v", request)
	return CreateOrderResponse{OrderID: 12345}, nil
}
