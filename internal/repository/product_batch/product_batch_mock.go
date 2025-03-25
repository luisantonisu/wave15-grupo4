package repsotory

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockProductBatchRepository struct {
	mock.Mock
}

func NewMockProductBatchRepository() *MockProductBatchRepository {
	return &MockProductBatchRepository{}
}

func (m *MockProductBatchRepository) Create(productBatch model.ProductBatchAttributes) (model.ProductBatch, error) {
	args := m.Called(productBatch)
	return args.Get(0).(model.ProductBatch), args.Error(1)
}
