package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetProduct() (map[int]model.Product, error) {
	args := m.Called()
	return args.Get(0).(map[int]model.Product), args.Error(1)
}

func (m *MockProductService) GetProductByID(id int) (model.Product, error) {
	args := m.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *MockProductService) GetProductRecord() (map[int]model.ProductRecordCount, error) {
	args := m.Called()
	return args.Get(0).(map[int]model.ProductRecordCount), args.Error(1)
}

func (m *MockProductService) GetProductRecordByID(id int) (model.ProductRecordCount, error) {
	args := m.Called(id)
	return args.Get(0).(model.ProductRecordCount), args.Error(1)
}

func (m *MockProductService) CreateProduct(product *model.ProductAttributes) (model.Product, error) {
	args := m.Called(product)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *MockProductService) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockProductService) UpdateProduct(id int, product *model.ProductAttributes) (*model.Product, error) {
	args := m.Called(id, product)
	return args.Get(0).(*model.Product), args.Error(1)
}
