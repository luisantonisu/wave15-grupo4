package service_test

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	carryRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/carry"
	localityRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/carry"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	carryID     = "C1"
	companyName = "Company 1"
	address     = "Address 1"
	telephone   = "123456789"
	LocalityID  = 1

	carry = model.Carry{
		ID: 1,
		CarryAttributes: model.CarryAttributes{
			CarryID:     &carryID,
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
			LocalityID:  &LocalityID,
		},
	}
)

func TestCarryService_Create(t *testing.T) {
	t.Run("case 1: create carry successfully", func(t *testing.T) {
		// Arrange
		carryRepo := carryRepository.NewCarryRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		carryService := service.NewCarryService(carryRepo, localityRepo)

		carryRepo.On("GetByCarryID", *carry.CarryID).Return(model.Carry{}, eh.GetErrNotFound(eh.CARRY_ID))
		localityRepo.On("GetByID", *carry.LocalityID).Return(model.LocalityDBModel{}, nil)
		carryRepo.On("Create", carry).Return(carry, nil)

		// Act
		result, err := carryService.Create(carry)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, carry, result)
		carryRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 2: invalid data - carry id empty", func(t *testing.T) {
		// Arrange
		carryRepo := carryRepository.NewCarryRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		carryService := service.NewCarryService(carryRepo, localityRepo)

		newCarry := carry
		newCarry.CarryID = nil

		// Act
		result, err := carryService.Create(newCarry)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, eh.GetErrInvalidData(service.ErrCarryIDEmpty), err)
		require.Equal(t, model.Carry{}, result)
		carryRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 3: conflict - carry id already exists", func(t *testing.T) {
		// Arrange
		carryRepo := carryRepository.NewCarryRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		carryService := service.NewCarryService(carryRepo, localityRepo)

		carryRepo.On("GetByCarryID", *carry.CarryID).Return(carry, nil)

		// Act
		result, err := carryService.Create(carry)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, eh.GetErrAlreadyExists(eh.CARRY_ID), err)
		require.Equal(t, model.Carry{}, result)
		carryRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 4: internal server error - get by carry id error", func(t *testing.T) {
		// Arrange
		carryRepo := carryRepository.NewCarryRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		carryService := service.NewCarryService(carryRepo, localityRepo)

		carryRepo.On("GetByCarryID", *carry.CarryID).Return(model.Carry{}, eh.GetErrDatabase(eh.CARRY))

		// Act
		result, err := carryService.Create(carry)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrDatabase)
		require.Equal(t, eh.GetErrDatabase(eh.CARRY), err)
		require.Equal(t, model.Carry{}, result)
		carryRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 5: foreign key - locality not found", func(t *testing.T) {
		// Arrange
		carryRepo := carryRepository.NewCarryRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		carryService := service.NewCarryService(carryRepo, localityRepo)

		carryRepo.On("GetByCarryID", *carry.CarryID).Return(model.Carry{}, eh.GetErrNotFound(eh.CARRY_ID))
		localityRepo.On("GetByID", *carry.LocalityID).Return(model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY))

		// Act
		result, err := carryService.Create(carry)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.LOCALITY), err)
		require.Equal(t, model.Carry{}, result)
		carryRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 6: database error - carry repository create", func(t *testing.T) {
		// Arrange
		carryRepo := carryRepository.NewCarryRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		carryService := service.NewCarryService(carryRepo, localityRepo)

		carryRepo.On("GetByCarryID", *carry.CarryID).Return(model.Carry{}, eh.GetErrNotFound(eh.CARRY_ID))
		localityRepo.On("GetByID", *carry.LocalityID).Return(model.LocalityDBModel{}, nil)
		carryRepo.On("Create", carry).Return(model.Carry{}, eh.GetErrDatabase(eh.CARRY))

		// Act
		result, err := carryService.Create(carry)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrDatabase)
		require.Equal(t, eh.GetErrDatabase(eh.CARRY), err)
		require.Equal(t, model.Carry{}, result)
		carryRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})
}
