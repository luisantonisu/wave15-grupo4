package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockSectionRepository struct {
	mock.Mock
}

func NewMockRepository() *MockSectionRepository {
	return &MockSectionRepository{}
}

func (m *MockSectionRepository) GetAll() ([]model.Section, error) {
	args := m.Called()
	return args.Get(0).([]model.Section), args.Error(1)
}

func (m *MockSectionRepository) GetByID(id int) (model.Section, error) {
	args := m.Called(id)
	return args.Get(0).(model.Section), args.Error(1)
}

func (m *MockSectionRepository) Create(section model.SectionAttributes) (model.Section, error) {
	args := m.Called(section)
	return args.Get(0).(model.Section), args.Error(1)
}

func (m *MockSectionRepository) Patch(id int, section model.SectionAttributes) (model.Section, error) {
	args := m.Called(id, section)
	return args.Get(0).(model.Section), args.Error(1)
}

func (m *MockSectionRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSectionRepository) Report(id *int) ([]model.ReportProductsBatches, error) {
	args := m.Called(id)
	return args.Get(0).([]model.ReportProductsBatches), args.Error(1)
}
