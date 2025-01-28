package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func MapEmployeeToEmployeeDTO(employees map[int]model.Employee) map[int]dto.EmployeeDTO {
	data := make(map[int]dto.EmployeeDTO)
	for _, value := range employees {
		data[value.Id] = dto.EmployeeDTO{
			Id:           value.Id,
			CardNumberId: value.CardNumberId,
			FirstName:    value.FirstName,
			LastName:     value.LastName,
			WarehouseId:  value.WarehouseId,
		}
	}
	return data
}
