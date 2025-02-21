package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

func NewProductMock() *ProductMock {
	return &ProductMock{}
}

type ProductMock struct {
	// FuncSearchProducts is the function that proxy the SearchProducts method.
	FuncGetProduct           func() (productMap map[int]model.Product, err error)
	FuncGetProductByID       func(id int) (product model.Product, err error)
	FuncGetProductRecord     func() (productRecordMap map[int]model.ProductRecordCount, err error)
	FuncGetProductRecordByID func(id int) (productRecord model.ProductRecordCount, err error)
	FuncCreateProduct        func(productAtrributes *model.ProductAttributes) (prod model.Product, err error)
	FuncDeleteProduct        func(id int) (err error)
	FuncUpdateProduct        func(id int, productAtrributes *model.ProductAttributes) (producto *model.Product, err error)
	// Spy
	Spy struct {
		// SearchProducts is the number of times the SearchProducts method is called.
		CreateProduct int
	}
}

/* type IProduct interface {
	GetProduct() (productMap map[int]model.Product, err error)
	GetProductByID(id int) (product model.Product, err error)
	GetProductRecord() (productRecordMap map[int]model.ProductRecordCount, err error)
	GetProductRecordByID(id int) (productRecord model.ProductRecordCount, err error)
	CreateProduct(productAtrributes *model.ProductAtrributes) (prod model.Product, err error)
	DeleteProduct(id int) (err error)
	UpdateProduct(id int, productAtrributes *model.ProductAtrributes) (producto *model.Product, err error)
} */
