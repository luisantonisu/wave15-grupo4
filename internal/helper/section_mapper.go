package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func SectionToSectionResponseDTO(section model.Section) dto.SectionResponseDTO {
	data := dto.SectionResponseDTO{
		Id:                 section.Id,
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseId:        section.WarehouseId,
		ProductTypeId:      section.ProductTypeId,
		ProductBatchId:     section.ProductBatchId,
	}
	return data
}

func SectionRequestDTOToSection(sectionRequestDTO dto.SectionRequestDTO) model.Section {
	data := model.Section{
		SectionAttributes: model.SectionAttributes{
			SectionNumber:      sectionRequestDTO.SectionNumber,
			CurrentTemperature: sectionRequestDTO.CurrentTemperature,
			MinimumTemperature: sectionRequestDTO.MinimumTemperature,
			CurrentCapacity:    sectionRequestDTO.CurrentCapacity,
			MinimumCapacity:    sectionRequestDTO.MinimumCapacity,
			MaximumCapacity:    sectionRequestDTO.MaximumCapacity,
			WarehouseId:        sectionRequestDTO.WarehouseId,
			ProductTypeId:      sectionRequestDTO.ProductTypeId,
			ProductBatchId:     sectionRequestDTO.ProductBatchId,
		},
	}
	return data
}
