package service

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
)

func NewEmployeeService(rp repository.IEmployee) *EmployeeService {
	return &EmployeeService{rp: rp}
}

type EmployeeService struct {
	rp repository.IEmployee
}

func (h *EmployeeService) GetAll() (map[int]model.Employee, error) {
	return h.rp.GetAll()
}

func (h *EmployeeService) GetByID(id int) (model.Employee, error) {
	return h.rp.GetByID(id)
}

func (h *EmployeeService) Create(employee model.Employee) (model.Employee, error) {
	if employee.FirstName == "" || employee.LastName == "" || employee.CardNumberID <= 0 || employee.WarehouseID <= 0 {
		return model.Employee{}, errors.New("Invalid employee data")
	}
	return h.rp.Create(employee)
}

func (h *EmployeeService) Update(id int, employee model.EmployeeAttributesPtr) (model.Employee, error) {
	return h.rp.Update(id, employee)
}

func (h *EmployeeService) Delete(id int) error {
	return h.rp.Delete(id)
}
