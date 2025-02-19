package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"

	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewCarryRepository(db *sql.DB) *CarryRepository {
	return &CarryRepository{
		db: db,
	}
}

type CarryRepository struct {
	db *sql.DB
}

func (r *CarryRepository) Create(carry model.Carry) (model.Carry, error) {
	row, err := r.db.Exec(`INSERT INTO carriers (
		carry_id, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)`,
		carry.CarryID, carry.CompanyName, carry.Address, carry.Telephone, carry.LocalityID)
	if err != nil {
		return model.Carry{}, eh.GetErrInvalidData(eh.CARRY)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.Carry{}, eh.GetErrDatabase(eh.CARRY)
	}

	return model.Carry{
		ID:              int(id),
		CarryAttributes: carry.CarryAttributes,
	}, nil
}

func (r *CarryRepository) GetByID(id int) (model.Carry, error) {
	var carry model.Carry
	err := r.db.QueryRow("SELECT id, carry_id, company_name, address, telephone, locality_id FROM carriers WHERE id = ?", id).Scan(
		&carry.ID, &carry.CarryID, &carry.CompanyName, &carry.Address, &carry.Telephone, &carry.LocalityID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Carry{}, eh.GetErrNotFound(eh.CARRY)
		}
		return model.Carry{}, eh.GetErrDatabase(eh.CARRY)
	}
	return carry, nil
}

func (r *CarryRepository) GetByCarryID(id string) (model.Carry, error) {
	var carry model.Carry
	err := r.db.QueryRow("SELECT id, carry_id, company_name, address, telephone, locality_id FROM carriers WHERE carry_id = ?", id).Scan(
		&carry.ID, &carry.CarryID, &carry.CompanyName, &carry.Address, &carry.Telephone, &carry.LocalityID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Carry{}, eh.GetErrNotFound(eh.CARRY)
		}
		return model.Carry{}, eh.GetErrDatabase(eh.CARRY)
	}
	return carry, nil
}
