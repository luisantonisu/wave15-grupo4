package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func EmployeeToEmployeeResponseDTO(employee model.Employee) dto.EmployeeResponseDTO {
	return dto.EmployeeResponseDTO{
		ID:           employee.ID,
		CardNumberID: employee.CardNumberID,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		WarehouseID:  employee.WarehouseID,
	}
}

func EmployeeRequestDTOToEmployee(employee dto.EmployeeRequestDTO) model.Employee {
	return model.Employee{
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: employee.CardNumberID,
			FirstName:    employee.FirstName,
			LastName:     employee.LastName,
			WarehouseID:  employee.WarehouseID,
		},
	}
}

func EmployeeResponseDTOToEmployee(employee dto.EmployeeResponseDTO) model.Employee {
	return model.Employee{
		ID: employee.ID,
		EmployeeAttributes: model.EmployeeAttributes{
			CardNumberID: employee.CardNumberID,
			FirstName:    employee.FirstName,
			LastName:     employee.LastName,
			WarehouseID:  employee.WarehouseID,
		},
	}
}

func EmployeeRequestDTOPtrToEmployeePtr(employee dto.EmployeeRequestDTOPtr) model.EmployeeAttributesPtr {
	return model.EmployeeAttributesPtr{
		CardNumberID: employee.CardNumberID,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		WarehouseID:  employee.WarehouseID,
	}
}
