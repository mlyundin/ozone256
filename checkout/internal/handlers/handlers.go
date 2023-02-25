package handlers

import "errors"

var (
	ErrEmptyUser = errors.New("empty user")
	ErrEmptySKU  = errors.New("empty sku")
)
