package deletefromcart

import (
	"context"
	"errors"
	"log"
	"route256/checkout/internal/domain"
)

type Handler struct {
	businessLogic *domain.Model
}

func New(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

type Request struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
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
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	log.Printf("deleteFromCart: %+v", req)

	// TODO add название и цена тянутся из ProductService.get_product
	var response Response

	return response, nil
}
