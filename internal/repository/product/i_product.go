package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IProduct interface {
	GetProduct() (productMap map[int]model.Product, err error)
	GetProductByID(id int) (product model.Product, err error)
	GetProductRecord() (productRecordMap map[int]model.ProductRecordCount, err error)
	GetProductRecordByID(id int) (productRecord model.ProductRecordCount, err error)
	CreateProduct(productAtrributes *model.ProductAtrributes) (prod model.Product, err error)
	DeleteProduct(id int) (err error)
	UpdateProduct(id int, productAtrributes *model.ProductAtrributesPtr) (producto *model.Product, err error)
}
