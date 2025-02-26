package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type WarehouseServiceMock struct {
	mock.Mock
}

func NewWarehouseServiceMock() *WarehouseServiceMock {
	return &WarehouseServiceMock{}
}

func (m *WarehouseServiceMock) GetAll() ([]model.Warehouse, error) {
	args := m.Called()
	return args.Get(0).([]model.Warehouse), args.Error(1)
}

func (m *WarehouseServiceMock) GetByID(id int) (model.Warehouse, error) {
	args := m.Called(id)
	return args.Get(0).(model.Warehouse), args.Error(1)
}

func (m *WarehouseServiceMock) Create(warehouse model.Warehouse) (model.Warehouse, error) {
	args := m.Called(warehouse)
	return args.Get(0).(model.Warehouse), args.Error(1)
}

func (m *WarehouseServiceMock) Update(id int, warehouse model.WarehouseAttributes) (model.Warehouse, error) {
	args := m.Called(id, warehouse)
	return args.Get(0).(model.Warehouse), args.Error(1)
}

func (m *WarehouseServiceMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
