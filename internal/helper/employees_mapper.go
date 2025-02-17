package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func EmployeeToEmployeeResponseDTO(employee model.Employee) dto.EmployeeResponseDTO {
	return dto.EmployeeResponseDTO{
		ID: employee.ID,
		EmployeeRequestDTO: dto.EmployeeRequestDTO{
			CardNumberID: employee.CardNumberID,
			FirstName:    employee.FirstName,
			LastName:     employee.LastName,
			WarehouseID:  employee.WarehouseID,
		},
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

func EmployeeRequestDTOPtrToEmployeePtr(employee dto.EmployeeRequestDTO) model.EmployeeAttributes {
	return model.EmployeeAttributes{
		CardNumberID: employee.CardNumberID,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		WarehouseID:  employee.WarehouseID,
	}
}

func InboundOrderToInboundOrderDTO(employees model.InboundOrdersReport) dto.InboundOrdersReportDTO {
	return dto.InboundOrdersReportDTO{
		ID:                 employees.ID,
		CardNumberID:       employees.CardNumberID,
		FirstName:          employees.FirstName,
		LastName:           employees.LastName,
		WarehouseID:        employees.WarehouseID,
		InboundOrdersCount: employees.InboundOrdersCount,
	}
}
