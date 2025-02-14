package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repositoryProductBatch "github.com/luisantonisu/wave15-grupo4/internal/repository/product_batch"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewProductBatchService(repositoryProductBatch repositoryProductBatch.IProductBatch) *ProductBatchService {
	return &ProductBatchService{
		productBatchRepo: repositoryProductBatch,
	}
}

type ProductBatchService struct {
	productBatchRepo repositoryProductBatch.IProductBatch
}

type ProductBatch struct {
	ID int
	ProductBatchAttributes
}

type ProductBatchAttributes struct {
	BatchNumber        int
	CurrentQuantity    int
	CurrentTemperature float64
	DueDate            string
	InitialQuantity    int
	ManufacturingDate  string
	ManufacturingHour  string
	MinimumTemperature float64
	ProductID          int
	SectionID          int
}

func (h *ProductBatchService) Create(productBatch model.ProductBatch) (model.ProductBatch, error) {
	if productBatch.BatchNumber <= 0 ||
		productBatch.CurrentQuantity <= 0 ||
		productBatch.CurrentTemperature <= 0 ||
		productBatch.DueDate == "" ||
		productBatch.InitialQuantity <= 0 ||
		productBatch.ManufacturingDate == "" ||
		productBatch.ManufacturingHour == "" ||
		productBatch.MinimumTemperature <= 0 ||
		productBatch.ProductID <= 0 ||
		productBatch.SectionID <= 0 {
		return model.ProductBatch{}, eh.GetErrInvalidData(eh.SECTION)
	}

	return h.productBatchRepo.Create(productBatch.ProductBatchAttributes)
}
