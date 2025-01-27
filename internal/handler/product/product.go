package handler

import (
	service "github.com/luisantonisu/wave15-grupo4/internal/service/product"
)

func NewProductHandler(sv service.ProductService) *ProductHandler {
	return &ProductHandler{sv: sv}
}

type ProductHandler struct {
	sv service.ProductService
}