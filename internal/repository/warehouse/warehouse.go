package repository

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func NewWarehouseRepository(db map[int]model.Warehouse) *WarehouseRepository {
	defaultDb := make(map[int]model.Warehouse)
	if db != nil {
		defaultDb = db
	}
	return &WarehouseRepository{db: defaultDb}
}

type WarehouseRepository struct {
	db map[int]model.Warehouse
}

func (wr *WarehouseRepository) GetAll() (map[int]model.Warehouse, error) {
	return wr.db, nil
}

func (wr *WarehouseRepository) GetByID(id int) (model.Warehouse, error) {
	if len(wr.db) == 0 {
		return model.Warehouse{}, errors.New("no warehouses found")
	}

	warehouse, ok := wr.db[id]
	if !ok {
		return model.Warehouse{}, errors.New("warehouse not found")
	}

	return warehouse, nil
}