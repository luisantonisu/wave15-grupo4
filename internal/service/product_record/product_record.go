package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/product_record"
)

type ProductRecordService struct {
	repository repository.IProductRecord
}

func NewProductRecordService(repository repository.IProductRecord) *ProductRecordService {
	return &ProductRecordService{repository: repository}
}

func (productRecordService *ProductRecordService) GetProductRecord() (map[int]model.ProductRecord, error) {
	return productRecordService.repository.GetProductRecord()
}

func (productRecordService *ProductRecordService) GetProductRecordByID(id int) (model.ProductRecord, error) {
	return productRecordService.repository.GetProductRecordByID(id)
}

func (productRecordService *ProductRecordService) CreateProductRecord(productRecord model.ProductRecordAtrributes) error {
	return productRecordService.repository.CreateProductRecord(productRecord)
}
