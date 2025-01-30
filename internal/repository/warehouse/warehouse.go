package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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
		return model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE)
	}

	warehouse, ok := wr.db[id]
	if !ok {
		return model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE)
	}

	return warehouse, nil
}

func (wr *WarehouseRepository) Create(warehouse model.Warehouse) (model.Warehouse, error) {
	if warehouse.WarehouseCode == "" {
		return model.Warehouse{}, eh.GetErrInvalidData(eh.WAREHOUSE_CODE)
	}

	if wr.warehouseCodeExists(warehouse.WarehouseCode) {
		return model.Warehouse{}, eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE)
	}

	wr.autoIncrement++
	warehouse.ID = wr.autoIncrement
	wr.db[warehouse.ID] = warehouse

	return warehouse, nil
}

func (wr *WarehouseRepository) Update(id int, warehouse model.Warehouse) (model.Warehouse, error) {
	updatedWarehouse, ok := wr.db[id]
	if !ok {
		return model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE)
	}

	if warehouse.WarehouseCode != "" && warehouse.WarehouseCode != updatedWarehouse.WarehouseCode {
		if wr.warehouseCodeExists(warehouse.WarehouseCode) {
			return model.Warehouse{}, eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE)
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

func (wr *WarehouseRepository) Delete(id int) error {
	_, ok := wr.db[id]
	if !ok {
		return eh.GetErrNotFound(eh.WAREHOUSE)
	}

	delete(wr.db, id)

	return nil
}

func (wr *WarehouseRepository) warehouseCodeExists(warehouseCode string) bool {
	for _, w := range wr.db {
		if w.WarehouseCode == warehouseCode {
			return true
		}
	}
	return false
}
