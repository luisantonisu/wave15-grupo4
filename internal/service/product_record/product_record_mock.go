package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockProductRecordService struct {
	mock.Mock
}

func NewMockReportService() *MockProductRecordService {
	return &MockProductRecordService{}
}

func (m *MockProductRecordService) CreateProductRecord(productRecord model.ProductRecordAtrributes) (prodRecord model.ProductRecord, err error) {
	args := m.Called()
	return args.Get(0).(model.ProductRecord), args.Error(1)
}
