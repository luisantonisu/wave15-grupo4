package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockProductBatchService struct {
	mock.Mock
}

func NewMockProductBatchService() *MockProductBatchService {
	return &MockProductBatchService{}
}

func (m *MockProductBatchService) Create(productBatch model.ProductBatchAttributes) (model.ProductBatch, error) {
	args := m.Called(productBatch)
	return args.Get(0).(model.ProductBatch), args.Error(1)
}
