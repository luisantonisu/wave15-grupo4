package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeRepository struct {
	mock.Mock
}

func NewEmployeeMockRepository() *MockEmployeeRepository {
	return &MockEmployeeRepository{}
}

func (m *MockEmployeeRepository) GetAll() (map[int]model.Employee, error) {
	args := m.Called()
	return args.Get(0).(map[int]model.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) GetByID(id int) (model.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) Create(employee model.EmployeeAttributes) (model.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEmployeeRepository) Update(id int, employee model.EmployeeAttributes) (model.Employee, error) {
	args := m.Called(id, employee)
	return args.Get(0).(model.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) Report(id int) (map[int]model.InboundOrdersReport, error) {
	args := m.Called(id)
	return args.Get(0).(map[int]model.InboundOrdersReport), args.Error(1)
}