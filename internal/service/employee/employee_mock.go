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

func (m *MockEmployeeService) GetEmployee() (map[int]model.Employee, error) {
	args := m.Called()
	return args.Get(0).(map[int]model.Employee), args.Error(1)
}

func (m *MockEmployeeService) GetEmployeeByID(id int) (*model.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Employee), args.Error(1)
}

func (m *MockEmployeeService) CreateEmployee(employee *model.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func (m *MockEmployeeService) DeleteEmployee(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEmployeeService) UpdateEmployee(id int, employee *model.Employee) (*model.Employee, error ){
	args := m.Called(id, employee)
	return args.Get(0).(*model.Employee), args.Error(0)
}



