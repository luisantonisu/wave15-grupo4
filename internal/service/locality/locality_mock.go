package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type LocalityServiceMock struct {
	mock.Mock
}

func NewLocalityServiceMock() *LocalityServiceMock {
	return &LocalityServiceMock{}
}

func (m *LocalityServiceMock) Create(locality model.Locality) (model.Locality, error) {
	args := m.Called(locality)
	return args.Get(0).(model.Locality), args.Error(1)
}

func (m *LocalityServiceMock) CarriersReport(id *int) ([]model.CarriersReport, error) {
	args := m.Called(id)
	return args.Get(0).([]model.CarriersReport), args.Error(1)
}

func (m *LocalityServiceMock) SellersReport(id *int) ([]model.LocalityReport, error) {
	args := m.Called(id)
	return args.Get(0).([]model.LocalityReport), args.Error(1)
}
