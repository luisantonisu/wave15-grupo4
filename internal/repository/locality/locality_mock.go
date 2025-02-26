package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

func NewLocalityRepositoryMock() *MockLocalityRepository {
	return &MockLocalityRepository{}
}

type MockLocalityRepository struct {
	mock.Mock
}

func (m *MockLocalityRepository) GetByID(id int) (model.LocalityDBModel, error) {
	args := m.Called(id)
	return args.Get(0).(model.LocalityDBModel), args.Error(1)
}

func (m *MockLocalityRepository) Create(locality model.LocalityDBModel) (model.LocalityDBModel, error) {
	args := m.Called(locality)
	return args.Get(0).(model.LocalityDBModel), args.Error(1)
}

func (m *MockLocalityRepository) CarriersReport(id *int) ([]model.CarriersReport, error) {
	args := m.Called(id)
	return args.Get(0).([]model.CarriersReport), args.Error(1)
}

func (m *MockLocalityRepository) SellersReport(id *int) ([]model.LocalityReport, error) {
	args := m.Called(id)
	return args.Get(0).([]model.LocalityReport), args.Error(1)
}
