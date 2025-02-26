package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type WarehouseRepositoryMock struct {
	mock.Mock
}

func NewWarehouseRepositoryMock() *WarehouseRepositoryMock {
	return &WarehouseRepositoryMock{}
}

func (m *WarehouseRepositoryMock) GetAll() ([]model.Warehouse, error) {
	args := m.Called()
	return args.Get(0).([]model.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) GetByID(id int) (model.Warehouse, error) {
	args := m.Called(id)
	return args.Get(0).(model.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) GetByCode(code string) (model.Warehouse, error) {
	args := m.Called(code)
	return args.Get(0).(model.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Create(warehouse model.Warehouse) (model.Warehouse, error) {
	args := m.Called(warehouse)
	return args.Get(0).(model.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Update(id int, warehouse model.Warehouse) (model.Warehouse, error) {
	args := m.Called(id, warehouse)
	return args.Get(0).(model.Warehouse), args.Error(1)
}

func (m *WarehouseRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
