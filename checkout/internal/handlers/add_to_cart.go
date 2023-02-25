package handlers

import (
	"context"
	"log"
	"route256/checkout/internal/domain"
)

type Handler struct {
	businessLogic *domain.Model
}

func NewAddToCartHandler(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

type AddToCartRequest struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (r AddToCartRequest) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}
	return nil
}

type AddToCartResponse struct {
}

func (h *Handler) Handle(ctx context.Context, req AddToCartRequest) (AddToCartResponse, error) {
	log.Printf("addToCart: %+v", req)

	var response AddToCartResponse

	err := h.businessLogic.AddToCart(ctx, req.User, req.Sku, req.Count)
	if err != nil {
		return response, err
	}

	return response, nil
}
