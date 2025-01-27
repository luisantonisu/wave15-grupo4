package repository

import (
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
