package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
)

func NewWarehouseService(rp repository.IWarehouse) *WarehouseService {
	return &WarehouseService{rp: rp}
}

type WarehouseService struct {
	rp repository.IWarehouse
}

func (ws *WarehouseService) GetAll() (map[int]model.Warehouse, error) {
	return ws.rp.GetAll()
}

func (ws *WarehouseService) GetByID(id int) (model.Warehouse, error) {
	return ws.rp.GetByID(id)
}
