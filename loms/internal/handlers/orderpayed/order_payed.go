package orderpayed

import (
	"context"
	"errors"
	"log"
)

type Request struct {
	OrderID int64 `json:"orderID"`
}

var (
	ErrEmptyOrder = errors.New("empty order")
)

func (r Request) Validate() error {
	if r.OrderID == 0 {
		return ErrEmptyOrder
	}
	return nil
}

type Response struct {
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("orderPayed: %+v", request)
	return Response{}, nil
}
