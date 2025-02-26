package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type BuyerServiceMock struct {
	mock.Mock
}

func NewBuyerServiceMock() *BuyerServiceMock {
	return &BuyerServiceMock{}
}

func (m *BuyerServiceMock) Create(buyer model.BuyerAttributes) (model.Buyer, error) {
	args := m.Called(buyer)
	return args.Get(0).(model.Buyer), args.Error(1)
}

func (m *BuyerServiceMock) GetAll() ([]model.Buyer, error) {
	args := m.Called()
	return args.Get(0).([]model.Buyer), args.Error(1)
}

func (m *BuyerServiceMock) GetByID(id int) (model.Buyer, error) {
	args := m.Called(id)
	return args.Get(0).(model.Buyer), args.Error(1)
}

func (m *BuyerServiceMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *BuyerServiceMock) Update(id int, buyer model.BuyerAttributes) (model.Buyer, error) {
	args := m.Called(id, buyer)
	return args.Get(0).(model.Buyer), args.Error(1)
}

func (m *BuyerServiceMock) PurchaseOrderReport(id *int) ([]model.ReportPurchaseOrders, error) {
	args := m.Called(id)
	return args.Get(0).([]model.ReportPurchaseOrders), args.Error(1)
}
