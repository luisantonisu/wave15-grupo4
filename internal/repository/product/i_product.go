package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IProduct interface {
	GetProduct() (productMap map[int]model.Product, err error)
	GetProductById(id int) (product model.Product, err error)
	CreateProduct(productAtrributes *model.ProductAtrributes) (err error)
}
