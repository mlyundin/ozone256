package purchase

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
	User int64 `json:"user"`
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
	log.Printf("purchase: %+v", req)

	var response Response
	// TODO Оформить заказ по всем товарам корзины. Вызывает createOrder у LOMS.

	return response, nil
}
