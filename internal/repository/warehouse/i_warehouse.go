package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IWarehouse interface {
	GetAll() (map[int]model.Warehouse, error)
	GetByID(id int) (model.Warehouse, error)
}
