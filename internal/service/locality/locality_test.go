package service_test

import (
	"errors"
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

var (
	localityID   = "1"
	localityName = "Merida"
	provinceName = "Yucatan"
	countryName  = "Mexico"

	locality1 = model.Locality{
		Id: localityID,
		LocalityAttributes: model.LocalityAttributes{
			LocalityName: &localityName,
			ProvinceName: &provinceName,
			CountryName:  &countryName,
		},
	}

	report1 = model.LocalityReport{
		Id:           1,
		LocalityName: "Callao",
		SellerCount:  10,
	}
	report2 = model.LocalityReport{
		Id:           1,
		LocalityName: "Merida",
		SellerCount:  9,
	}

	SellerReport = []model.LocalityReport{report1, report2}
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

func TestLocalityService_SellersReport(t *testing.T) {
	t.Run("case 1: get sellers report successfully", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		localityRepo.On("SellersReport", mock.Anything).Return(SellerReport, nil)

		// Act
		result, err := localityService.SellersReport(nil)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, SellerReport, result)
		countryRepository.AssertExpectations(t)
		provinceRepository.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})
	t.Run("case 2: get carriers sellers successfully with locality id", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		localityRepo.On("SellersReport", mock.Anything).Return([]model.LocalityReport{report1}, nil)

		// Act
		result, err := localityService.SellersReport(&report1.Id)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, []model.LocalityReport{report1}, result)
		countryRepository.AssertExpectations(t)
		provinceRepository.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})
	t.Run("case 3 : Error found", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		localityRepo.On("SellersReport", mock.Anything).Return([]model.LocalityReport{}, errors.New("error test"))

		//Act
		result, err := localityService.SellersReport(nil)

		expectedResult := []model.LocalityReport{}

		//Assert
		require.Error(t, err)
		require.Equal(t, "error test", err.Error())
		require.Equal(t, expectedResult, result)
		countryRepository.AssertExpectations(t)
		provinceRepository.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})
}

func TestLocalityService_ValidateLocality(t *testing.T) {
	locality := model.Locality{
		Id: localityID,
		LocalityAttributes: model.LocalityAttributes{
			LocalityName: &localityName,
			ProvinceName: &provinceName,
			CountryName:  &countryName,
		},
	}
	t.Run("case 1: validate is successfully", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()

		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)
		
		//Act
		err := localityService.ValidateLocality(locality)

		//Assert
		require.NoError(t, err)
	})

	t.Run("case 2: LocalityId is required", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		locality.Id = ""

		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)
		
		//Act
		err := localityService.ValidateLocality(locality)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, "invalid data: locality", err.Error())
	})
	t.Run("case 3: LocalityId should contain only numbers", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		locality.Id = "Aaa"

		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)
		
		//Act
		err := localityService.ValidateLocality(locality)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, "invalid data: locality", err.Error())
	})
	t.Run("case 4: LocalityName is required", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		locality.Id = localityID
		locality.LocalityName = new(string)

		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)
		
		//Act
		err := localityService.ValidateLocality(locality)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, "invalid data: locality", err.Error())
	})
	t.Run("case 5: ProvinceName is required", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		locality.LocalityName = &localityName
		locality.ProvinceName = new(string)

		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)
		
		//Act
		err := localityService.ValidateLocality(locality)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, "invalid data: locality", err.Error())
	})
	t.Run("case 6: CountryName is required", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		locality.ProvinceName = &provinceName
		locality.CountryName = new(string)

		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)
		
		//Act
		err := localityService.ValidateLocality(locality)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, "invalid data: locality", err.Error())
	})
}

func TestLocalityService_Create(t *testing.T) {
	localityDb := model.LocalityDBModel{
		Id:           1,
		LocalityName: localityName,
		ProvinceID:   1,
	}
	locality := model.Locality{
		Id: "1",
		LocalityAttributes: model.LocalityAttributes{
			LocalityName: &localityName,
			ProvinceName: &provinceName,
			CountryName:  &countryName,
		},
	}
	InvalidLocality := model.Locality{
		Id: "1",
		LocalityAttributes: model.LocalityAttributes{
			LocalityName: new(string),
			ProvinceName: new(string),
			CountryName:  new(string),
		},
	}

	t.Run("case 1: create locality successfully", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		countryRepository.On("GetCountryIDByCountryName", "Mexico").Return(1, nil)
		provinceRepository.On("GetProvinceID", 1, "Yucatan").Return(1, nil)
		localityRepo.On("Create", localityDb).Return(localityDb, nil)

		//Act
		result, err := localityService.Create(locality1)

		//Assert
		require.NoError(t, err)
		require.NotNil(t, localityDb)
		require.Equal(t, locality, result)
		countryRepository.AssertExpectations(t)
		provinceRepository.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})
	t.Run("case 2: Invalid data", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		//Act
		result, err := localityService.Create(InvalidLocality)

		//Assert
		require.Error(t, err)
		require.Equal(t, model.Locality{}, result)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, "invalid data: locality", err.Error())
	})
	t.Run("case 3: Invalid country name", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		countryRepository.On("GetCountryIDByCountryName", "Mexico").Return(0, errors.New("test error"))

		//Act
		result, err := localityService.Create(locality1)

		//Assert
		require.Error(t, err)
		require.Equal(t, model.Locality{}, result)
		require.Equal(t, "test error", err.Error())
	})
	t.Run("case 4: Invalid province name", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		countryRepository.On("GetCountryIDByCountryName", "Mexico").Return(1, nil)
		provinceRepository.On("GetProvinceID", 1, "Yucatan").Return(0, errors.New("test error"))

		//Act
		result, err := localityService.Create(locality1)

		//Assert
		require.Error(t, err)
		require.Equal(t, model.Locality{}, result)
		require.Equal(t, "test error", err.Error())
	})
	t.Run("case 5: Error in Create Locality", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		countryRepository.On("GetCountryIDByCountryName", "Mexico").Return(1, nil)
		provinceRepository.On("GetProvinceID", 1, "Yucatan").Return(1, nil)
		localityRepo.On("Create", localityDb).Return(model.LocalityDBModel{}, errors.New("test error"))

		//Act
		result, err := localityService.Create(locality1)

		//Assert
		require.Error(t, err)
		require.Equal(t, model.Locality{}, result)
		require.Equal(t, "test error", err.Error())
	})
	t.Run("case 6: localityID value out of rang", func(t *testing.T) {
		// Arrange
		countryRepository := countryRepository.NewCountryRepositoryMock()
		provinceRepository := provinceRepository.NewProvinceRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		localityService := service.NewLocalityService(countryRepository, provinceRepository, localityRepo)

		countryRepository.On("GetCountryIDByCountryName", "Mexico").Return(1, nil)
		provinceRepository.On("GetProvinceID", 1, "Yucatan").Return(1, nil)

		locality1.Id = "9999999999999999999" 

		localityRepo.On("Create", localityDb).Return(model.LocalityDBModel{}, errors.New("test error"))

		//Act
		result, err := localityService.Create(locality1)

		//Assert
		require.Error(t, err)
		require.Equal(t, model.Locality{}, result)
		require.Equal(t, "strconv.Atoi: parsing \"9999999999999999999\": value out of range", err.Error())
	})
}
