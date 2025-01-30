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
	return &WarehouseRepository{autoIncrement: len(db), db: defaultDb}
}

type WarehouseRepository struct {
	autoIncrement int
	db            map[int]model.Warehouse
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

func (wr *WarehouseRepository) Create(warehouse model.Warehouse) (model.Warehouse, error) {
	if warehouse.WarehouseCode == "" {
		return model.Warehouse{}, errors.New("warehouse code is required")
	}

	if wr.warehouseCodeExists(warehouse.WarehouseCode) {
		return model.Warehouse{}, errors.New("warehouse code already exists")
	}

	wr.autoIncrement++
	warehouse.ID = wr.autoIncrement
	wr.db[warehouse.ID] = warehouse

	return warehouse, nil
}

func (wr *WarehouseRepository) Update(id int, warehouse model.Warehouse) (model.Warehouse, error) {
	updatedWarehouse, ok := wr.db[id]
	if !ok {
		return model.Warehouse{}, errors.New("warehouse not found")
	}

	if warehouse.WarehouseCode != "" && warehouse.WarehouseCode != updatedWarehouse.WarehouseCode {
		if wr.warehouseCodeExists(warehouse.WarehouseCode) {
			return model.Warehouse{}, errors.New("warehouse code already exists")
		}
		updatedWarehouse.WarehouseCode = warehouse.WarehouseCode
	}

	if warehouse.Address != "" {
		updatedWarehouse.Address = warehouse.Address
	}

	if warehouse.Telephone > 0 {
		updatedWarehouse.Telephone = warehouse.Telephone
	}

	if warehouse.MinimumCapacity > 0 {
		updatedWarehouse.MinimumCapacity = warehouse.MinimumCapacity
	}

	if warehouse.MinimumTemperature > -50 && warehouse.MinimumTemperature < 40 {
		updatedWarehouse.MinimumTemperature = warehouse.MinimumTemperature
	}

	wr.db[id] = updatedWarehouse

	return updatedWarehouse, nil
}

func (wr *WarehouseRepository) warehouseCodeExists(warehouseCode string) bool {
	for _, w := range wr.db {
		if w.WarehouseCode == warehouseCode {
			return true
		}
	}
	return false
}
