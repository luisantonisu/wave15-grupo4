package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

/* func NewProductMock() *ProductMock {
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
} */

/* type IProduct interface {
	GetProduct() (productMap map[int]model.Product, err error)
	GetProductByID(id int) (product model.Product, err error)
	GetProductRecord() (productRecordMap map[int]model.ProductRecordCount, err error)
	GetProductRecordByID(id int) (productRecord model.ProductRecordCount, err error)
	CreateProduct(productAtrributes *model.ProductAtrributes) (prod model.Product, err error)
	DeleteProduct(id int) (err error)
	UpdateProduct(id int, productAtrributes *model.ProductAtrributes) (producto *model.Product, err error)
} */

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetProduct() (map[int]model.Product, error) {
	args := m.Called()
	return args.Get(0).(map[int]model.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductByID(id int) (model.Product, error) {
	args := m.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *MockProductRepository) GetProductRecord() (map[int]model.ProductRecordCount, error) {
	args := m.Called()
	return args.Get(0).(map[int]model.ProductRecordCount), args.Error(1)
}

func (m *MockProductRepository) GetProductRecordByID(id int) (model.ProductRecordCount, error) {
	args := m.Called(id)
	return args.Get(0).(model.ProductRecordCount), args.Error(1)
}

func (m *MockProductRepository) CreateProduct(product *model.ProductAttributes) (model.Product, error) {
	args := m.Called(product)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductRepository) UpdateProduct(id int, product *model.ProductAttributes) (*model.Product, error) {
	args := m.Called(id, product)
	return args.Get(0).(*model.Product), args.Error(1)
}
