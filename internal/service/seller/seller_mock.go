package service


import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

func NewSellerServiceMock() *MockSellerService {
	return &MockSellerService{}
}

type MockSellerService struct {
	mock.Mock
}

func (m *MockSellerService) GetAll() (sellers []model.Seller, err error) {
	args := m.Called()
	return args.Get(0).([]model.Seller), args.Error(1)
}

func (m *MockSellerService) GetByID(id int) (seller model.Seller, err error) {
	args := m.Called(id)
	return args.Get(0).(model.Seller), args.Error(1)
}

func (m *MockSellerService) Create(seller model.SellerAttributes) (model.Seller, error) {
	args := m.Called(seller)
	return args.Get(0).(model.Seller), args.Error(1)
}

func (m *MockSellerService) Update(id int, seller model.SellerAttributes) (model.Seller, error) {
	args := m.Called(id, seller)
	return args.Get(0).(model.Seller), args.Error(1)
}

func (m *MockSellerService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}



