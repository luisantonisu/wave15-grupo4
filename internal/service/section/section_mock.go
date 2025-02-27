package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockSectionService struct {
	mock.Mock
}

func NewSectionMock() *MockSectionService {
	return &MockSectionService{}
}

func (m *MockSectionService) GetAll() ([]model.Section, error) {
	args := m.Called()
	return args.Get(0).([]model.Section), args.Error(1)
}

func (m *MockSectionService) GetByID(id int) (model.Section, error) {
	args := m.Called(id)
	return args.Get(0).(model.Section), args.Error(1)
}

func (m *MockSectionService) Create(section model.SectionAttributes) (model.Section, error) {
	args := m.Called(section)
	return args.Get(0).(model.Section), args.Error(1)
}

func (m *MockSectionService) Patch(id int, section model.SectionAttributes) (model.Section, error) {
	args := m.Called(id, section)
	return args.Get(0).(model.Section), args.Error(1)
}

func (m *MockSectionService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSectionService) Report(id *int) ([]model.ReportProductsBatches, error) {
	args := m.Called(id)
	return args.Get(0).([]model.ReportProductsBatches), args.Error(1)
}
