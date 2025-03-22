package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockProductRecordRepository struct {
	mock.Mock
}

func NewMockRepository() *MockProductRecordRepository {
	return &MockProductRecordRepository{}
}

func (m *MockProductRecordRepository) CreateProductRecord(productRecord model.ProductRecordAtrributes) (err error) {
	args := m.Called()
	return args.Error(0)
}
