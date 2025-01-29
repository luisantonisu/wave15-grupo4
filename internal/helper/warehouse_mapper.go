package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func WarehouseToWarehouseResponseDTO(warehouse model.Warehouse) dto.WarehouseResponseDTO {
	data := dto.WarehouseResponseDTO{
		Id:                 warehouse.Id,
		WarehouseCode:      warehouse.WarehouseCode,
		Address:            warehouse.Address,
		Telephone:          warehouse.Telephone,
		MinimumCapacity:    warehouse.MinimumCapacity,
		MinimumTemperature: warehouse.MinimumTemperature,
	}
	return data
}

func WarehouseRequestDTOToWarehouse(warehouseRequestDTO dto.WarehouseRequestDTO) model.Warehouse {
	data := model.Warehouse{
		WarehouseAttributes: model.WarehouseAttributes{
			WarehouseCode:      warehouseRequestDTO.WarehouseCode,
			Address:            warehouseRequestDTO.Address,
			Telephone:          warehouseRequestDTO.Telephone,
			MinimumCapacity:    warehouseRequestDTO.MinimumCapacity,
			MinimumTemperature: warehouseRequestDTO.MinimumTemperature,
		},
	}
	return data
}
