package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type InboundOrderRepoMock struct {
	mock.Mock
}

func NewInboundOrderRepoMock() *InboundOrderRepoMock {
	return &InboundOrderRepoMock{}
}

func (m *InboundOrderRepoMock) AlreadyExists(attribute string, value int) bool {
	args := m.Called(attribute, value)
	return args.Get(0).(bool)
}

func (m *InboundOrderRepoMock) CreateInboundOrder(inboundOrder model.InboundOrderAttributes) (model.InboundOrder, error) {
	args := m.Called(inboundOrder)
	return args.Get(0).(model.InboundOrder), args.Error(1)
}
