package repository

import (
	"database/sql"
	"fmt"

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
	fmt.Println("carry ", carry)
	row, err := r.db.Exec(`INSERT INTO carriers (
		carry_id, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)`,
		carry.CarryID, carry.CompanyName, carry.Address, carry.Telephone, carry.LocalityID)
	if err != nil {
		fmt.Println(err)
		return model.Carry{}, eh.GetErrInvalidData(eh.CARRY)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.Carry{}, eh.GetErrInvalidData(eh.CARRY)
	}

	var newCarry model.Carry
	newCarry.ID = int(id)
	newCarry.CarryAttributes = carry.CarryAttributes

	return newCarry, nil
}

func (r *CarryRepository) GetByID(id int) (model.Carry, error) {
		// Create response entity
		var carry model.Carry

		// Get carry from db
		err := r.db.QueryRow("SELECT id, company_id, company_name, address, telephone, locality_id FROM carriers WHERE id = ?", id).Scan(
			&carry.ID, &carry.CarryID, &carry.CompanyName, &carry.Address, &carry.Telephone, &carry.LocalityID,
		)
		if err != nil {
			return model.Carry{}, eh.GetErrNotFound(eh.CARRY)
		}
		return carry, nil
}
