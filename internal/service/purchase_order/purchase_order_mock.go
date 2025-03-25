package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type PurchaseOrderServiceMock struct {
	mock.Mock
}

func NewPurchaseOrderServiceMock() *PurchaseOrderServiceMock {
	return &PurchaseOrderServiceMock{}
}

func (m *PurchaseOrderServiceMock) Create(purchaseOrder model.PurchaseOrderAttributes) (model.PurchaseOrder, error) {
	args := m.Called(purchaseOrder)
	return args.Get(0).(model.PurchaseOrder), args.Error(1)
}
