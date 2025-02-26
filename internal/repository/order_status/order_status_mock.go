package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type OrderStatusRepositoryMock struct {
	mock.Mock
}

func NewOrderStatusRepositoryMock() *OrderStatusRepositoryMock {
	return &OrderStatusRepositoryMock{}
}

func (m *OrderStatusRepositoryMock) GetByID(id int) (model.OrderStatus, error) {
	args := m.Called(id)
	return args.Get(0).(model.OrderStatus), args.Error(1)
}
