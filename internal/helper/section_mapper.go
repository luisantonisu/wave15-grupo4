package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func SectionToSectionResponseDTO(section model.Section) dto.SectionResponseDTO {
	return dto.SectionResponseDTO{
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
}

func SectionRequestDTOToSection(section dto.SectionRequestDTO) model.Section {
	return model.Section{
		SectionAttributes: model.SectionAttributes{
			SectionNumber:      section.SectionNumber,
			CurrentTemperature: section.CurrentTemperature,
			MinimumTemperature: section.MinimumTemperature,
			CurrentCapacity:    section.CurrentCapacity,
			MinimumCapacity:    section.MinimumCapacity,
			MaximumCapacity:    section.MaximumCapacity,
			WarehouseID:        section.WarehouseID,
			ProductTypeID:      section.ProductTypeID,
			ProductBatchID:     section.ProductBatchID,
		},
	}
}

func SectionResponseDTOToSection(section dto.SectionResponseDTO) model.Section {
	return model.Section{
		ID: section.ID,
		SectionAttributes: model.SectionAttributes{
			SectionNumber:      section.SectionNumber,
			CurrentTemperature: section.CurrentTemperature,
			MinimumTemperature: section.MinimumTemperature,
			CurrentCapacity:    section.CurrentCapacity,
			MinimumCapacity:    section.MinimumCapacity,
			MaximumCapacity:    section.MaximumCapacity,
			WarehouseID:        section.WarehouseID,
			ProductTypeID:      section.ProductTypeID,
			ProductBatchID:     section.ProductBatchID,
		},
	}
}

func SectionRequestDTOPtrToSectionPtr(section dto.SectionRequestDTOPtr) model.SectionAttributesPtr {
	return model.SectionAttributesPtr{
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
}
