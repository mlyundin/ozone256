package handlers

import (
	"context"
	"log"
	"route256/checkout/internal/domain"
)

type DeleteFromCartHandler struct {
	businessLogic *domain.Model
}

func NewDeleteFromCartHandler(businessLogic *domain.Model) *DeleteFromCartHandler {
	return &DeleteFromCartHandler{
		businessLogic: businessLogic,
	}
}

type DeleteFromCartRequest struct {
	User  int64  `json:"user"`
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (r DeleteFromCartRequest) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}
	return nil
}

type DeleteFromCartResponse struct {
}

func (h *DeleteFromCartHandler) Handle(ctx context.Context, req DeleteFromCartRequest) (DeleteFromCartResponse, error) {
	log.Printf("deleteFromCart: %+v", req)

	// TODO add название и цена тянутся из ProductService.get_product
	var response DeleteFromCartResponse

	return response, nil
}
