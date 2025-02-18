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

func (r *WarehouseRepository) GetAll() ([]model.Warehouse, error) {
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

	warehouses := []model.Warehouse{}
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
		warehouses = append(warehouses, warehouse)
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

func (r *WarehouseRepository) GetByCode(code string) (model.Warehouse, error) {
	var warehouse model.Warehouse
	err := r.db.QueryRow(`SELECT id, 
			warehouse_code, 
			address, 
			telephone, 
			minimum_capacity, 
			minimum_temperature,
			locality_id
			FROM warehouses WHERE warehouse_code = ?`, code).
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
		return model.Warehouse{}, eh.GetErrDatabase(eh.WAREHOUSE)
	}

	return model.Warehouse{
		ID:                  int(id),
		WarehouseAttributes: warehouse.WarehouseAttributes,
	}, nil
}

func (r *WarehouseRepository) Update(id int, warehouse model.Warehouse) (model.Warehouse, error) {
	_, err := r.db.Exec(`UPDATE warehouses SET
		warehouse_code = ?, address = ?, telephone = ?, minimum_capacity = ?, minimum_temperature = ?, locality_id = ? WHERE id = ?`,
		warehouse.WarehouseCode, warehouse.Address, warehouse.Telephone,
		warehouse.MinimumCapacity, warehouse.MinimumTemperature, warehouse.LocalityID, id)
	if err != nil {
		return model.Warehouse{}, eh.GetErrInvalidData(eh.WAREHOUSE)
	}

	return warehouse, nil
}

func (r *WarehouseRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM warehouses WHERE id = ?", id)
	if err != nil {
		return eh.GetErrDatabase(eh.WAREHOUSE)
	}

	return nil
}
