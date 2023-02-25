package handlers

import (
	"errors"
)

var (
	ErrEmptyOrder = errors.New("empty order")
)

type Item struct {
	Sku   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func NewHandler[Handler any]() *Handler {
	var h Handler
	return &h
}
