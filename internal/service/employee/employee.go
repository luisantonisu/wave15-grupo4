package service

import (
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
)

func NewEmployeeService(rp repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{rp: rp}
}

type EmployeeService struct {
	rp repository.EmployeeRepository
}

