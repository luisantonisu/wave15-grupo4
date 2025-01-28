package handler

import (
	service "github.com/luisantonisu/wave15-grupo4/internal/service/employee"
)

func NewEmployeeHandler(sv service.IEmployee) *EmployeeHandler {
	return &EmployeeHandler{sv: sv}
}

type EmployeeHandler struct {
	sv service.IEmployee
}