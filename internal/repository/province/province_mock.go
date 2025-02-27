package repository

import (
	"github.com/stretchr/testify/mock"
)

type ProvinceRepositoryMock struct {
	mock.Mock
}

func NewProvinceRepositoryMock() *ProvinceRepositoryMock {
	return &ProvinceRepositoryMock{}
}

func (m *ProvinceRepositoryMock) GetProvinceID(countryID int, provinceName string) (int, error) {
	args := m.Called(countryID, provinceName)
	return args.Get(0).(int), args.Error(1)
}
