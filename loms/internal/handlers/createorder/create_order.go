package createorder

import (
	"context"
	"errors"
	"log"
)

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type Request struct {
	User  int64  `json:"user"`
	Items []Item `json:"items"`
}

var (
	ErrEmptyUser = errors.New("empty user")
)

func (r Request) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}
	return nil
}

type Response struct {
	OrderID int64 `json:"orderID"`
}

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, request Request) (Response, error) {
	log.Printf("createOrder: %+v", request)
	return Response{OrderID: 12345}, nil
}
