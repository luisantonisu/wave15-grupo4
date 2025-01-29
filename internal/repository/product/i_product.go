package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IProduct interface {
	GetProduct() (productMap map[int]model.Product, err error)
	GetProductByID(id int) (product model.Product, err error)
	CreateProduct(productAtrributes *model.ProductAtrributes) (err error)
	DeleteProduct(id int) (err error)
	UpdateProduct(id int, productAtrributes *model.ProductAtrributes) (producto *model.Product, err error)
}
