package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type CarryServiceMock struct {
	mock.Mock
}

func NewCarryServiceMock() *CarryServiceMock {
	return &CarryServiceMock{}
}

func (m *CarryServiceMock) Create(carry model.Carry) (model.Carry, error) {
	args := m.Called(carry)
	return args.Get(0).(model.Carry), args.Error(1)
}
