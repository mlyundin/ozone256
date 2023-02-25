package handlers

import (
	"errors"
	"route256/checkout/internal/domain"
)

var (
	ErrEmptyUser = errors.New("empty user")
	ErrEmptySKU  = errors.New("empty sku")
)

type Handler struct {
	businessLogic *domain.Model
}

func NewHandler(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}
