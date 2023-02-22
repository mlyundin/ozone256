package listorder

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

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Response struct {
	Status string `json:"status"`
	User   int64  `json:"user"`
	Items  []Item `json:"items"`
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("listOrder: %+v", request)
	return Response{
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
