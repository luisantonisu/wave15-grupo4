package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IWarehouse interface {
	GetAll() ([]model.Warehouse, error)
	GetByID(id int) (model.Warehouse, error)
	Create(warehouse model.Warehouse) (model.Warehouse, error)
	Update(id int, warehouse model.WarehouseAttributes) (model.Warehouse, error)
	Delete(id int) error
}
