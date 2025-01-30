package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func SectionToSectionResponseDTO(section model.Section) dto.SectionResponseDTO {
	data := dto.SectionResponseDTO{
		ID:                 section.ID,
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseID:        section.WarehouseID,
		ProductTypeID:      section.ProductTypeID,
		ProductBatchID:     section.ProductBatchID,
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
			WarehouseID:        sectionRequestDTO.WarehouseID,
			ProductTypeID:      sectionRequestDTO.ProductTypeID,
			ProductBatchID:     sectionRequestDTO.ProductBatchID,
		},
	}
	return data
}
