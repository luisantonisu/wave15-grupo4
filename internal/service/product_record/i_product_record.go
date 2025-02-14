package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IProductRecord interface {
	CreateProductRecord(productRecord model.ProductRecordAtrributes) error
}
