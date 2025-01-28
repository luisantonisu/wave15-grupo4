package service

import (
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
)

func NewEmployeeService(rp repository.IEmployee) *EmployeeService {
	return &EmployeeService{rp: rp}
}

type EmployeeService struct {
	rp repository.IEmployee
}

