package handlers

import (
	"context"
	"log"
)

type PurchaseRequest struct {
	User int64 `json:"user"`
}

func (r PurchaseRequest) Validate() error {
	if r.User == 0 {
		return ErrEmptyUser
	}
	return nil
}

type PurchaseResponse struct {
}

func (h *Handler) Purchase(ctx context.Context, req PurchaseRequest) (PurchaseResponse, error) {
	log.Printf("purchase: %+v", req)

	var response PurchaseResponse

	err := h.businessLogic.Purchase(ctx, req.User)
	if err != nil {
		return response, err
	}

	return response, nil
}
