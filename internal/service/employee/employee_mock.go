package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeService struct {
	mock.Mock
}

func NewEmployeeMock() *MockEmployeeService {
	return &MockEmployeeService{}
}

// Create implements IEmployee.
func (m *MockEmployeeService) Create(employee model.Employee) (model.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(model.Employee), args.Error(1)
}

// Delete implements IEmployee.
func (m *MockEmployeeService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetAll implements IEmployee.
func (m *MockEmployeeService) GetAll() (map[int]model.Employee, error) {
	args := m.Called()
	return args.Get(0).(map[int]model.Employee), args.Error(1)
}

// GetByID implements IEmployee.
func (m *MockEmployeeService) GetByID(id int) (model.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(model.Employee), args.Error(1)
}

// Report implements IEmployee.
func (m *MockEmployeeService) Report(id int) (map[int]model.InboundOrdersReport, error) {
	args := m.Called(id)
	return args.Get(0).(map[int]model.InboundOrdersReport), args.Error(1)
}

// Update implements IEmployee.
func (m *MockEmployeeService) Update(id int, employee model.EmployeeAttributes) (model.Employee, error) {
	args := m.Called(id, employee)
	return args.Get(0).(model.Employee), args.Error(1)
}
