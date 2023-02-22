package cancelorder

import (
	"context"
	"log"
)

type Request struct {
	OrderID int64 `json:"orderID"`
}

func (r Request) Validate() error {
	return nil
}

type Response struct {
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("cancelOrder: %+v", request)
	return Response{}, nil
}
