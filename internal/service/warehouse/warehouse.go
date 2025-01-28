package service

import (
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
)

func NewWarehouseService(rp repository.IWarehouse) *WarehouseService {
	return &WarehouseService{rp: rp}
}

type WarehouseService struct {
	rp repository.IWarehouse
}