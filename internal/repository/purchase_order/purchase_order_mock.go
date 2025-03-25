package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type PurchaseOrderRepositoryMock struct {
	mock.Mock
}

func NewPurchaseOrderRepositoryMock() *PurchaseOrderRepositoryMock {
	return &PurchaseOrderRepositoryMock{}
}

func (m *PurchaseOrderRepositoryMock) Create(purchaseOrder model.PurchaseOrderAttributes) (model.PurchaseOrder, error) {
	args := m.Called(purchaseOrder)
	return args.Get(0).(model.PurchaseOrder), args.Error(1)
}
func (m *PurchaseOrderRepositoryMock) OrderNumberExists(orderNumber string) bool {
	args := m.Called(orderNumber)
	return args.Bool(0)
}
