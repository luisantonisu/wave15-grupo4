package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type InboundOrderServiceMock struct {
	mock.Mock
}

func NewInboundOrderServiceMock() *InboundOrderServiceMock {
	return &InboundOrderServiceMock{}
}

func (m *InboundOrderServiceMock) Create(inboundOrder model.InboundOrderAttributes) (model.InboundOrder, error) {
	args := m.Called()
	return args.Get(0).(model.InboundOrder), args.Error(1)
}
