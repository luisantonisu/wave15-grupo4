package repository

import (
	"github.com/stretchr/testify/mock"
)

type CountryRepositoryMock struct {
	mock.Mock
}

func NewCountryRepositoryMock() *CountryRepositoryMock {
	return &CountryRepositoryMock{}
}

func (m *CountryRepositoryMock) GetCountryIDByCountryName(countryName string) (int, error) {
	args := m.Called(countryName)
	return args.Get(0).(int), args.Error(1)
}
