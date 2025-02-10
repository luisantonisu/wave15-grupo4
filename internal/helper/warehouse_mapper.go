package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func WarehouseToWarehouseResponseDTO(warehouse model.Warehouse) dto.WarehouseResponseDTO {
	data := dto.WarehouseResponseDTO{
		ID:                 warehouse.ID,
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

func WarehouseRequestDTOPtrToWarehouseAttributesPtr(warehouseRequestDTOPtr dto.WarehouseRequestDTOPtr) model.WarehouseAttributesPtr {
	data := model.WarehouseAttributesPtr{
			WarehouseCode:      warehouseRequestDTOPtr.WarehouseCode,
			Address:            warehouseRequestDTOPtr.Address,
			Telephone:          warehouseRequestDTOPtr.Telephone,
			MinimumCapacity:    warehouseRequestDTOPtr.MinimumCapacity,
			MinimumTemperature: warehouseRequestDTOPtr.MinimumTemperature,
		}
	return data
}
