package handler

import (
	service "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
)

func NewBuyerHandler(sv service.BuyerService) *BuyerHandler {
	return &BuyerHandler{sv: sv}
}

type BuyerHandler struct {
	sv service.BuyerService
}