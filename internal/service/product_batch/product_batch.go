package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repositoryProduct "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	repositoryProductBatch "github.com/luisantonisu/wave15-grupo4/internal/repository/product_batch"
	repositorySection "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewProductBatchService(repositoryProductBatch repositoryProductBatch.IProductBatch, repositorySection repositorySection.ISection, repositoryProduct repositoryProduct.IProduct) *ProductBatchService {
	return &ProductBatchService{
		productBatchRepo: repositoryProductBatch,
		productRepo:      repositoryProduct,
		sectionRepo:      repositorySection,
	}
}

type ProductBatchService struct {
	productBatchRepo repositoryProductBatch.IProductBatch
	productRepo      repositoryProduct.IProduct
	sectionRepo      repositorySection.ISection
}

func (h *ProductBatchService) Create(productBatch model.ProductBatchAttributes) (model.ProductBatch, error) {
	// Validate if the product exists
	_, err := h.productRepo.GetProductByID(productBatch.ProductID)
	if err != nil {
		return model.ProductBatch{}, eh.GetErrForeignKey(eh.PRODUCT)
	}

	// Validate if the section exists
	_, err = h.sectionRepo.GetByID(productBatch.SectionID)
	if err != nil {
		return model.ProductBatch{}, eh.GetErrForeignKey(eh.SECTION)
	}

	return h.productBatchRepo.Create(productBatch)
}
