package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func MapSectionToSectionDTO(sections map[int]model.Section) map[int]dto.SectionDTO {
	data := make(map[int]dto.SectionDTO)
	for _, value := range sections {
		data[value.Id] = dto.SectionDTO{
			Id:                 value.Id,
			SectionNumber:      value.SectionNumber,
			CurrentTemperature: value.CurrentTemperature,
			MinimumTemperature: value.MinimumTemperature,
			CurrentCapacity:    value.CurrentCapacity,
			MinimumCapacity:    value.MinimumCapacity,
			MaximumCapacity:    value.MaximumCapacity,
			WarehouseId:        value.WarehouseId,
			ProductTypeId:      value.ProductTypeId,
			ProductBatchId:     value.ProductBatchId,
		}
	}
	return data
}
