package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockProductRecordRepository struct {
	mock.Mock
}

func NewMockReportRepository() *MockProductRecordRepository {
	return &MockProductRecordRepository{}
}

func (m *MockProductRecordRepository) CreateProductRecord(productRecord model.ProductRecordAtrributes) (prodRecord model.ProductRecord, err error) {
	args := m.Called()
	return args.Get(0).(model.ProductRecord), args.Error(1)
}
