package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type BuyerRepositoryMock struct {
	mock.Mock
}

func NewBuyerRepositoryMock() *BuyerRepositoryMock {
	return &BuyerRepositoryMock{}
}

func (m *BuyerRepositoryMock) Create(buyer model.BuyerAttributes) (model.Buyer, error) {
	args := m.Called(buyer)
	return args.Get(0).(model.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) GetAll() ([]model.Buyer, error) {
	args := m.Called()
	return args.Get(0).([]model.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) GetByID(id int) (model.Buyer, error) {
	args := m.Called(id)
	return args.Get(0).(model.Buyer), args.Error(1)
}
func (m *BuyerRepositoryMock) GetByCardNumberID(id string) (model.Buyer, error) {
	args := m.Called(id)
	return args.Get(0).(model.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *BuyerRepositoryMock) Update(id int, buyer model.BuyerAttributes) (model.Buyer, error) {
	args := m.Called(id, buyer)
	return args.Get(0).(model.Buyer), args.Error(1)
}

func (m *BuyerRepositoryMock) PurchaseOrderReport(id *int) ([]model.ReportPurchaseOrders, error) {
	args := m.Called(id)
	return args.Get(0).([]model.ReportPurchaseOrders), args.Error(1)
}
