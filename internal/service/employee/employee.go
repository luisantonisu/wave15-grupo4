package service

import (
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
