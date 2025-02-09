package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

type IProductRecord interface {
	GetProductRecord() (productRecordMap map[int]model.ProductRecord, err error)
	GetProductRecordByID(id int) (productRecord model.ProductRecord, err error)
	CreateProductRecord(productRecord model.ProductRecordAtrributes) (err error)
}
