package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IEmployee interface {
	CardNumberIDExists(int, int) bool
	EmployeeExists(int) bool
	GetAll() ([]model.Employee, error)
	GetByID(id int) (model.Employee, error)
	Create(employee model.EmployeeAttributes) (model.Employee, error)
	Update(id int, employee model.EmployeeAttributes) (model.Employee, error)
	Delete(id int) error
	Report(id int) ([]model.InboundOrdersReport, error)
}
