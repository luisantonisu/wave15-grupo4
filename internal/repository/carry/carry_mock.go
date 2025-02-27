package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type CarryRepositoryMock struct {
	mock.Mock
}

func NewCarryRepositoryMock() *CarryRepositoryMock {
	return &CarryRepositoryMock{}
}

func (m *CarryRepositoryMock) Create(carry model.Carry) (model.Carry, error) {
	args := m.Called(carry)
	return args.Get(0).(model.Carry), args.Error(1)
}

func (m *CarryRepositoryMock) GetByID(id int) (model.Carry, error) {
	args := m.Called(id)
	return args.Get(0).(model.Carry), args.Error(1)
}

func (m *CarryRepositoryMock) GetByCarryID(id string) (model.Carry, error) {
	args := m.Called(id)
	return args.Get(0).(model.Carry), args.Error(1)
}
