package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repositoryProductBatch "github.com/luisantonisu/wave15-grupo4/internal/repository/product_batch"
)

func NewProductBatchService(repositoryProductBatch repositoryProductBatch.IProductBatch) *ProductBatchService {
	return &ProductBatchService{
		productBatchRepo: repositoryProductBatch,
	}
}

type ProductBatchService struct {
	productBatchRepo repositoryProductBatch.IProductBatch
}

func (h *ProductBatchService) Create(productBatch model.ProductBatchAttributes) (model.ProductBatch, error) {
	// if productBatch.BatchNumber == "" ||
	// 	productBatch.CurrentQuantity <= 0 ||
	// 	productBatch.CurrentTemperature <= 0 ||
	// 	productBatch.DueDate == "" ||
	// 	productBatch.InitialQuantity <= 0 ||
	// 	productBatch.ManufacturingDate == "" ||
	// 	productBatch.ManufacturingHour == "" ||
	// 	productBatch.MinimumTemperature <= 0 ||
	// 	productBatch.ProductID <= 0 ||
	// 	productBatch.SectionID <= 0 {
	// 	return model.ProductBatch{}, eh.GetErrInvalidData(eh.SECTION)
	// }

	return h.productBatchRepo.Create(productBatch)
}
