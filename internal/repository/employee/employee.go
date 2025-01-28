package repository

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func NewEmployeeRepository(db map[int]model.Employee) *EmployeeRepository {
	defaultDb := make(map[int]model.Employee)
	if db != nil {
		defaultDb = db
	}
	return &EmployeeRepository{db: defaultDb}
}

type EmployeeRepository struct {
	db map[int]model.Employee
}

func (r *EmployeeRepository) employeeExists(id int) bool {
	_, exists := r.db[id]
	return exists
}

func (r *EmployeeRepository) GetAll() (map[int]model.Employee, error) {
	return r.db, nil
}

func (r *EmployeeRepository) GetByID(id int) (model.Employee, error) {
	if !r.employeeExists(id) {
		return model.Employee{}, errors.New("employee not found")
	}

	return r.db[id], nil
}
