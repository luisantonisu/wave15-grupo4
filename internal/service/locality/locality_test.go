package service_test

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	countryRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/country"
	localityRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"
	provinceRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/province"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/locality"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLocalityService_CarriersReport(t *testing.T) {
	record1 := model.CarriersReport{
		LocalityID:    1,
		LocalityName:  "Locality 1",
		CarriersCount: 10,
	}
	record2 := model.CarriersReport{
		LocalityID:    2,
		LocalityName:  "Locality 2",
		CarriersCount: 1,
	}
	report := []model.CarriersReport{record1, record2}

	t.Run("case 1: get carriers report successfully", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		localityRepo.On("CarriersReport", mock.Anything).Return(report, nil)

		// Act
		result, err := localityService.CarriersReport(nil)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, report, result)
		countryRepository.AssertExpectations(t)
		provinceRepository.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 2: get carriers report successfully with locality id", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		localityRepo.On("GetByID", mock.Anything).Return(model.LocalityDBModel{}, nil)
		localityRepo.On("CarriersReport", mock.Anything).Return([]model.CarriersReport{record1}, nil)

		// Act
		result, err := localityService.CarriersReport(&record1.LocalityID)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, []model.CarriersReport{record1}, result)
		countryRepository.AssertExpectations(t)
		provinceRepository.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 3: not found - locality id not found", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		localityRepo.On("GetByID", mock.Anything).Return(model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY))

		// Act
		result, err := localityService.CarriersReport(&record1.LocalityID)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, eh.GetErrNotFound(eh.LOCALITY), err)
		require.Equal(t, []model.CarriersReport(nil), result)
		countryRepository.AssertExpectations(t)
		provinceRepository.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})
}
