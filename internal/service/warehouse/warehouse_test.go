package service_test

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	localityRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	warehouseCodeA = "WH1"
	warehouseCodeB = "WH2"
	address        = "Address 1"
	telephone      = uint(123456789)
	capacity       = 100
	temperature    = float32(20.5)
	localityID     = 1
	warehouseA     = model.Warehouse{
		ID: 1,
		WarehouseAttributes: model.WarehouseAttributes{
			WarehouseCode:      &warehouseCodeA,
			Address:            &address,
			Telephone:          &telephone,
			MinimumCapacity:    &capacity,
			MinimumTemperature: &temperature,
			LocalityID:         &localityID,
		},
	}
	warehouseB = model.Warehouse{
		ID: 2,
		WarehouseAttributes: model.WarehouseAttributes{
			WarehouseCode:      &warehouseCodeB,
			Address:            &address,
			Telephone:          &telephone,
			MinimumCapacity:    &capacity,
			MinimumTemperature: &temperature,
			LocalityID:         &localityID,
		},
	}
)

func TestWarehouseService_GetAll(t *testing.T) {
	t.Run("case 1: get all warehouses successfully", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)
		warehouses := []model.Warehouse{warehouseA, warehouseB}

		warehouseRepo.On("GetAll").Return(warehouses, nil)

		// Act
		result, err := warehouseService.GetAll()

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, warehouses, result)
		warehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: no warehouses found", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		warehouseRepo.On("GetAll").Return([]model.Warehouse{}, nil)

		// Act
		result, err := warehouseService.GetAll()

		// Assert
		require.NoError(t, err)
		require.Empty(t, result)
		require.Equal(t, []model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
	})
}

func TestWarehouseService_GetByID(t *testing.T) {
	t.Run("case 1: get warehouse by ID successfully", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		warehouseRepo.On("GetByID", 1).Return(warehouseA, nil)

		// Act
		result, err := warehouseService.GetByID(1)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, warehouseA, result)
		warehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: not found - warehouse not found", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errNotFound := eh.GetErrNotFound(eh.WAREHOUSE)
		warehouseRepo.On("GetByID", 1).Return(model.Warehouse{}, errNotFound)

		// Act
		result, err := warehouseService.GetByID(1)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errNotFound, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
	})
}

func TestWarehouseService_Create(t *testing.T) {
	t.Run("case 1: create warehouse successfully", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)
		warehouseRepo.On("GetByCode", *warehouseA.WarehouseCode).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE_CODE))
		localityRepo.On("GetByID", *warehouseA.LocalityID).Return(model.LocalityDBModel{Id: *warehouseA.LocalityID}, nil)
		warehouseRepo.On("Create", warehouseA).Return(warehouseA, nil)

		// Act
		result, err := warehouseService.Create(warehouseA)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, warehouseA, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 2: invalid data - warehouse code empty", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errInvalidData := eh.GetErrInvalidData("warehouse code is empty")

		// Act
		result, err := warehouseService.Create(model.Warehouse{})

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 3: conflict - warehouse code already exists", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errConflict := eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE)
		warehouseRepo.On("GetByCode", *warehouseA.WarehouseCode).Return(warehouseA, nil)

		// Act
		result, err := warehouseService.Create(warehouseA)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, errConflict, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 4: internal server error - get by warehouse code error", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errInternalError := eh.GetErrDatabase(eh.WAREHOUSE)
		warehouseRepo.On("GetByCode", *warehouseA.WarehouseCode).Return(model.Warehouse{}, eh.GetErrDatabase(eh.WAREHOUSE))

		// Act
		result, err := warehouseService.Create(warehouseA)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrDatabase)
		require.Equal(t, errInternalError, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 5: invalid data - invalid telephone", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errInvalidData := eh.GetErrInvalidData(service.ErrTelephoneEmptyOrInvalid)
		warehouseRepo.On("GetByCode", *warehouseA.WarehouseCode).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE_CODE))

		badPhone := uint(123)
		badPhoneWarehouse := warehouseA
		badPhoneWarehouse.Telephone = &badPhone
		// Act
		result, err := warehouseService.Create(badPhoneWarehouse)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 6: invalid data - invalid minimum capacity", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errInvalidData := eh.GetErrInvalidData(service.ErrMinimumCapacityNegative)
		warehouseRepo.On("GetByCode", *warehouseA.WarehouseCode).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE_CODE))

		badMinCap := -1
		badMinCapWarehouse := warehouseA
		badMinCapWarehouse.MinimumCapacity = &badMinCap
		// Act
		result, err := warehouseService.Create(badMinCapWarehouse)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 7: invalid data - invalid minimum temperature", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errInvalidData := eh.GetErrInvalidData(service.ErrMinimumTemperatureOutOfRange)
		warehouseRepo.On("GetByCode", *warehouseA.WarehouseCode).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE_CODE))

		badMinTemp := float32(-100)
		badMinTempWarehouse := warehouseA
		badMinTempWarehouse.MinimumTemperature = &badMinTemp
		// Act
		result, err := warehouseService.Create(badMinTempWarehouse)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 8: foreign key - locality id doesn't exists", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errForeignKey := eh.GetErrForeignKey(eh.LOCALITY)
		warehouseRepo.On("GetByCode", *warehouseA.WarehouseCode).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE_CODE))
		localityRepo.On("GetByID", *warehouseA.LocalityID).Return(model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY))

		// Act
		result, err := warehouseService.Create(warehouseA)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, errForeignKey, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})
}

func TestWarehouseService_Update(t *testing.T) {
	t.Run("case 1: update warehouse successfully", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		warehouseRepo.On("GetByID", warehouseA.ID).Return(warehouseA, nil)
		localityRepo.On("GetByID", *warehouseA.LocalityID).Return(model.LocalityDBModel{Id: *warehouseA.LocalityID}, nil)
		warehouseRepo.On("Update", warehouseA.ID, warehouseA).Return(warehouseA, nil)

		// Act
		result, err := warehouseService.Update(warehouseA.ID, warehouseA.WarehouseAttributes)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, warehouseA, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 2: not found - warehouse not found", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errNotFound := eh.GetErrNotFound(eh.WAREHOUSE)
		warehouseRepo.On("GetByID", warehouseA.ID).Return(model.Warehouse{}, errNotFound)

		// Act
		result, err := warehouseService.Update(warehouseA.ID, warehouseA.WarehouseAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errNotFound, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 3: conflict - warehouse code already exists", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)
		newWarehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseB.WarehouseAttributes,
		}

		errConflict := eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE)
		warehouseRepo.On("GetByID", warehouseA.ID).Return(warehouseA, nil)
		warehouseRepo.On("GetByCode", *newWarehouse.WarehouseCode).Return(newWarehouse, nil)

		// Act
		result, err := warehouseService.Update(newWarehouse.ID, newWarehouse.WarehouseAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, errConflict, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 4: internal server error - get by warehouse code error", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)
		newWarehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseB.WarehouseAttributes,
		}

		errInternalError := eh.GetErrDatabase(eh.WAREHOUSE)
		warehouseRepo.On("GetByID", warehouseA.ID).Return(warehouseA, nil)
		warehouseRepo.On("GetByCode", *newWarehouse.WarehouseCode).Return(model.Warehouse{}, eh.GetErrDatabase(eh.WAREHOUSE))

		// Act
		result, err := warehouseService.Update(newWarehouse.ID, newWarehouse.WarehouseAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrDatabase)
		require.Equal(t, errInternalError, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 5: invalid data - invalid telephone", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errInvalidData := eh.GetErrInvalidData(service.ErrTelephoneEmptyOrInvalid)
		warehouseRepo.On("GetByID", warehouseA.ID).Return(warehouseA, nil)

		badPhone := uint(123)
		badPhoneWarehouse := warehouseA
		badPhoneWarehouse.Telephone = &badPhone
		// Act
		result, err := warehouseService.Update(badPhoneWarehouse.ID, badPhoneWarehouse.WarehouseAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 6: invalid data - invalid minimum capacity", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errInvalidData := eh.GetErrInvalidData(service.ErrMinimumCapacityNegative)
		warehouseRepo.On("GetByID", warehouseA.ID).Return(warehouseA, nil)

		badMinCap := -1
		badMinCapWarehouse := warehouseA
		badMinCapWarehouse.MinimumCapacity = &badMinCap
		// Act
		result, err := warehouseService.Update(badMinCapWarehouse.ID, badMinCapWarehouse.WarehouseAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 7: invalid data - invalid minimum temperature", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errInvalidData := eh.GetErrInvalidData(service.ErrMinimumTemperatureOutOfRange)
		warehouseRepo.On("GetByID", warehouseA.ID).Return(warehouseA, nil)

		badMinTemp := float32(-100)
		badMinTempWarehouse := warehouseA
		badMinTempWarehouse.MinimumTemperature = &badMinTemp
		// Act
		result, err := warehouseService.Update(badMinTempWarehouse.ID, badMinTempWarehouse.WarehouseAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 8: foreign key - locality id doesn't exists", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errForeignKey := eh.GetErrForeignKey(eh.LOCALITY)
		warehouseRepo.On("GetByID", warehouseA.ID).Return(warehouseA, nil)
		localityRepo.On("GetByID", *warehouseA.LocalityID).Return(model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY))

		// Act
		result, err := warehouseService.Update(warehouseA.ID, warehouseA.WarehouseAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, errForeignKey, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})
}

func TestWarehouseService_Delete(t *testing.T) {
	t.Run("case 1: delete warehouse successfully", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		warehouseRepo.On("GetByID", warehouseA.ID).Return(warehouseA, nil)
		warehouseRepo.On("Delete", warehouseA.ID).Return(nil)

		// Act
		err := warehouseService.Delete(warehouseA.ID)

		// Assert
		require.NoError(t, err)
		warehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: not found - warehouse not found", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)

		errNotFound := eh.GetErrNotFound(eh.WAREHOUSE)
		warehouseRepo.On("GetByID", 1).Return(model.Warehouse{}, errNotFound)

		// Act
		err := warehouseService.Delete(1)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errNotFound, err)
		warehouseRepo.AssertExpectations(t)
	})
}
