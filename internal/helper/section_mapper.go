package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func SectionToSectionResponseDTO(section model.Section) dto.SectionResponseDTO {
	data := dto.SectionResponseDTO{
		ID: section.ID,
		SectionRequestDTO: dto.SectionRequestDTO{
			SectionNumber:      section.SectionNumber,
			CurrentTemperature: section.CurrentTemperature,
			MinimumTemperature: section.MinimumTemperature,
			CurrentCapacity:    section.CurrentCapacity,
			MinimumCapacity:    section.MinimumCapacity,
			MaximumCapacity:    section.MaximumCapacity,
			WarehouseID:        section.WarehouseID,
			ProductTypeID:      section.ProductTypeID,
		},
	}

	return data
}

func SectionRequestDTOToSection(section dto.SectionRequestDTO) model.SectionAttributes {
	data := model.SectionAttributes{
		SectionNumber:      section.SectionNumber,
		CurrentTemperature: section.CurrentTemperature,
		MinimumTemperature: section.MinimumTemperature,
		CurrentCapacity:    section.CurrentCapacity,
		MinimumCapacity:    section.MinimumCapacity,
		MaximumCapacity:    section.MaximumCapacity,
		WarehouseID:        section.WarehouseID,
		ProductTypeID:      section.ProductTypeID,
	}
	return data
}

func SectionRequestDTOPtrToSectionPtr(sectionRequestDTO dto.SectionRequestDTO) model.SectionAttributes {
	data := model.SectionAttributes{
		SectionNumber:      sectionRequestDTO.SectionNumber,
		CurrentTemperature: sectionRequestDTO.CurrentTemperature,
		MinimumTemperature: sectionRequestDTO.MinimumTemperature,
		CurrentCapacity:    sectionRequestDTO.CurrentCapacity,
		MinimumCapacity:    sectionRequestDTO.MinimumCapacity,
		MaximumCapacity:    sectionRequestDTO.MaximumCapacity,
		WarehouseID:        sectionRequestDTO.WarehouseID,
		ProductTypeID:      sectionRequestDTO.ProductTypeID,
	}
	return data
}

func ReportProductsBatchesToReportProductsBatchesResponseDTO(report model.ReportProductsBatches) dto.ReportProductsBatchesResponseDTO {
	data := dto.ReportProductsBatchesResponseDTO{
		SectionID:     report.SectionID,
		SectionNumber: report.SectionNumber,
		ProductsCount: report.ProductsCount,
	}
	return data
}
