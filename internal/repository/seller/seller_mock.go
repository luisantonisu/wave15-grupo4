package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

func NewSellerRepositoryMock() *MockSellerRepository {
	return &MockSellerRepository{}
}

type MockSellerRepository struct {
	mock.Mock
}

func (m *MockSellerRepository) GetAll() ([]model.Seller, error) {
	args := m.Called()
	return args.Get(0).([]model.Seller), args.Error(1)
}

func (m *MockSellerRepository) GetByID(id int) (model.Seller, error) {
	args := m.Called(id)
	return args.Get(0).(model.Seller), args.Error(1)
}

func (m *MockSellerRepository) Create(seller model.SellerAttributes) (model.Seller, error) {
	args := m.Called(seller)
	return args.Get(0).(model.Seller), args.Error(1)
}

func (m *MockSellerRepository) Update(id int, seller model.SellerAttributes) (model.Seller, error) {
	args := m.Called(id, seller)
	return args.Get(0).(model.Seller), args.Error(1)
}

func (m *MockSellerRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}



