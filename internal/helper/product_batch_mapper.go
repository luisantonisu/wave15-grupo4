package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func ProductBatchToProductBatchResponseDTO(productBatch model.ProductBatch) dto.ProductBatchResponseDTO {
	return dto.ProductBatchResponseDTO{
		ID:                 productBatch.ID,
		BatchNumber:        productBatch.BatchNumber,
		CurrentQuantity:    productBatch.CurrentQuantity,
		CurrentTemperature: productBatch.CurrentTemperature,
		DueDate:            productBatch.DueDate,
		InitialQuantity:    productBatch.InitialQuantity,
		ManufacturingDate:  productBatch.ManufacturingDate,
		ManufacturingHour:  productBatch.ManufacturingHour,
		MinimumTemperature: productBatch.MinimumTemperature,
		ProductID:          productBatch.ProductID,
		SectionID:          productBatch.SectionID,
	}
}

func ProductBatchRequestDTOToProductBatch(productBatch dto.ProductBatchRequestDTO) model.ProductBatchAttributes {
	return model.ProductBatchAttributes{
		BatchNumber:        productBatch.BatchNumber,
		CurrentQuantity:    productBatch.CurrentQuantity,
		CurrentTemperature: productBatch.CurrentTemperature,
		DueDate:            productBatch.DueDate,
		InitialQuantity:    productBatch.InitialQuantity,
		ManufacturingDate:  productBatch.ManufacturingDate,
		ManufacturingHour:  productBatch.ManufacturingHour,
		MinimumTemperature: productBatch.MinimumTemperature,
		ProductID:          productBatch.ProductID,
		SectionID:          productBatch.SectionID,
	}
}
