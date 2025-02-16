package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewWarehouseRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{
		db: db,
	}
}

type WarehouseRepository struct {
	db *sql.DB
}

func (r *WarehouseRepository) GetAll() (map[int]model.Warehouse, error) {
	rows, err := r.db.Query(`SELECT id, 
		warehouse_code, 
		address, 
		telephone, 
		minimum_capacity, 
		minimum_temperature,
		locality_id FROM warehouses`)

	if err != nil {
		return nil, eh.GetErrGettingData(eh.WAREHOUSE)
	}
	defer rows.Close()

	warehouses := make(map[int]model.Warehouse)
	for rows.Next() {
		var warehouse model.Warehouse
		err := rows.Scan(&warehouse.ID,
			&warehouse.WarehouseCode,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.MinimumCapacity,
			&warehouse.MinimumTemperature,
			&warehouse.LocalityID)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.WAREHOUSE)
		}
		warehouses[warehouse.ID] = warehouse
	}
	return warehouses, nil
}

func (r *WarehouseRepository) GetByID(id int) (model.Warehouse, error) {
	var warehouse model.Warehouse
	err := r.db.QueryRow(`SELECT id, 
			warehouse_code, 
			address, 
			telephone, 
			minimum_capacity, 
			minimum_temperature,
			locality_id
			FROM warehouses WHERE id = ?`, id).
		Scan(&warehouse.ID,
			&warehouse.WarehouseCode,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.MinimumCapacity,
			&warehouse.MinimumTemperature,
			&warehouse.LocalityID)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE)
		}
		return model.Warehouse{}, eh.GetErrParsingData(eh.WAREHOUSE)
	}

	return warehouse, nil
}

func (r *WarehouseRepository) Create(warehouse model.Warehouse) (model.Warehouse, error) {
	row, err := r.db.Exec(`INSERT INTO warehouses (
		warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) VALUES (?, ?, ?, ?, ?, ?)`,
		warehouse.WarehouseCode, warehouse.Address, warehouse.Telephone, warehouse.MinimumCapacity, warehouse.MinimumTemperature, warehouse.LocalityID)
	if err != nil {
		return model.Warehouse{}, eh.GetErrInvalidData(eh.WAREHOUSE)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.Warehouse{}, eh.GetErrInvalidData(eh.WAREHOUSE)
	}

	var newWarehouse model.Warehouse
	newWarehouse.ID = int(id)
	newWarehouse.WarehouseAttributes = warehouse.WarehouseAttributes

	return newWarehouse, nil
}

func (r *WarehouseRepository) Update(id int, warehouse model.WarehouseAttributesPtr) (model.Warehouse, error) {
	existingWarehouse, err := r.GetByID(id)
	if err != nil {
		return model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE)
	}

	if warehouse.WarehouseCode != nil && *warehouse.WarehouseCode != existingWarehouse.WarehouseCode {
		existingWarehouse.WarehouseCode = *warehouse.WarehouseCode
	}

	if warehouse.Address != nil {
		existingWarehouse.Address = *warehouse.Address
	}

	if warehouse.Telephone != nil {
		existingWarehouse.Telephone = *warehouse.Telephone
	}

	if warehouse.MinimumCapacity != nil {
		existingWarehouse.MinimumCapacity = *warehouse.MinimumCapacity
	}

	if warehouse.MinimumTemperature != nil {
		existingWarehouse.MinimumTemperature = *warehouse.MinimumTemperature
	}

	_, err = r.db.Exec(`UPDATE warehouses SET
		warehouse_code = ?, address = ?, telephone = ?, minimum_capacity = ?, minimum_temperature = ?, locality_id = ? WHERE id = ?`,
		existingWarehouse.WarehouseCode, existingWarehouse.Address, existingWarehouse.Telephone,
		existingWarehouse.MinimumCapacity, existingWarehouse.MinimumTemperature, existingWarehouse.LocalityID, id)
	if err != nil {
		return model.Warehouse{}, eh.GetErrInvalidData(eh.WAREHOUSE)
	}

	return existingWarehouse, nil
}

func (r *WarehouseRepository) Delete(id int) error {
	_, err := r.GetByID(id)
	if err != nil {
		return eh.GetErrNotFound(eh.WAREHOUSE)
	}

	_, err = r.db.Exec("DELETE FROM warehouses WHERE id = ?", id)
	if err != nil {
		return eh.GetErrNotFound(eh.WAREHOUSE)
	}

	return nil
}
