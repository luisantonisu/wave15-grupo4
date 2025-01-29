package handler

import (
	service "github.com/luisantonisu/wave15-grupo4/internal/service/seller"
)

func NewSellerHandler(sv service.ISeller) *SellerHandler {
	return &SellerHandler{sv: sv}
}

type SellerHandler struct {
	sv service.ISeller
}