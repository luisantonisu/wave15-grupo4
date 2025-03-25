package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repoProduct "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/product_record"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type ProductRecordService struct {
	repository        repository.IProductRecord
	repositoryProduct repoProduct.IProduct
}

func NewProductRecordService(repositoryRecord repository.IProductRecord, repositoryProduct repoProduct.IProduct) *ProductRecordService {
	return &ProductRecordService{
		repository:        repositoryRecord,
		repositoryProduct: repositoryProduct,
	}
}

func (s *ProductRecordService) CreateProductRecord(productRecord model.ProductRecordAtrributes) (prodRecord model.ProductRecord, err error) {
	_, err = s.repositoryProduct.GetProductByID(*productRecord.ProductId)

	if err != nil {
		return model.ProductRecord{}, errorHandler.GetErrForeignKey(errorHandler.PRODUCT)
	}
	return s.repository.CreateProductRecord(productRecord)
}
