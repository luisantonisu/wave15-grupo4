package handler

import (
	service "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"
)

func NewWarehouseHandler(sv service.IWarehouse) *WarehouseHandler {
	return &WarehouseHandler{sv: sv}
}

type WarehouseHandler struct {
	sv service.IWarehouse
}