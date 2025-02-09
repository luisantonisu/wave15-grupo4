package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IProductRecord interface {
	GetProductRecord() (map[int]model.ProductRecord, error)
	GetProductRecordByID(id int) (model.ProductRecord, error)
	CreateProductRecord(productRecord model.ProductRecordAtrributes) error
}
