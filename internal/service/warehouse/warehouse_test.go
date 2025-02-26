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
	warehouseCodeA       = "WH1"
	warehouseCodeB       = "WH2"
	address              = "Address 1"
	telephone            = uint(123456789)
	capacity             = 100
	temperature          = float32(20.5)
	localityID           = 1
	warehouseAttributesA = model.WarehouseAttributes{
		WarehouseCode:      &warehouseCodeA,
		Address:            &address,
		Telephone:          &telephone,
		MinimumCapacity:    &capacity,
		MinimumTemperature: &temperature,
		LocalityID:         &localityID,
	}
	warehouseAttributesB = model.WarehouseAttributes{
		WarehouseCode:      &warehouseCodeB,
		Address:            &address,
		Telephone:          &telephone,
		MinimumCapacity:    &capacity,
		MinimumTemperature: &temperature,
		LocalityID:         &localityID,
	}
)

func TestWarehouseService_GetAll(t *testing.T) {
	t.Run("case 1: get all warehouses successfully", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)
		warehouses := []model.Warehouse{
			{
				ID:                  1,
				WarehouseAttributes: warehouseAttributesA,
			},
			{
				ID:                  2,
				WarehouseAttributes: warehouseAttributesB,
			},
		}

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
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}

		warehouseRepo.On("GetByID", 1).Return(warehouse, nil)

		// Act
		result, err := warehouseService.GetByID(1)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, warehouse, result)
		warehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: warehouse not found", func(t *testing.T) {
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
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}

		warehouseRepo.On("GetByCode", *warehouse.WarehouseCode).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE_CODE))
		localityRepo.On("GetByID", *warehouse.LocalityID).Return(model.LocalityDBModel{Id: *warehouse.LocalityID}, nil)
		warehouseRepo.On("Create", warehouse).Return(warehouse, nil)

		// Act
		result, err := warehouseService.Create(warehouse)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, warehouse, result)
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
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}

		errConflict := eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE)
		warehouseRepo.On("GetByCode", *warehouse.WarehouseCode).Return(warehouse, nil)

		// Act
		result, err := warehouseService.Create(warehouse)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, errConflict, err)
		require.Equal(t, model.Warehouse{}, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 4: foreign key - locality id doesn't exists", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}

		errForeignKey := eh.GetErrForeignKey(eh.LOCALITY)
		warehouseRepo.On("GetByCode", *warehouse.WarehouseCode).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE_CODE))
		localityRepo.On("GetByID", *warehouse.LocalityID).Return(model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY))

		// Act
		result, err := warehouseService.Create(warehouse)

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
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}

		warehouseRepo.On("GetByID", warehouse.ID).Return(warehouse, nil)
		localityRepo.On("GetByID", *warehouse.LocalityID).Return(model.LocalityDBModel{Id: *warehouse.LocalityID}, nil)
		warehouseRepo.On("Update", warehouse.ID, warehouse).Return(warehouse, nil)

		// Act
		result, err := warehouseService.Update(warehouse.ID, warehouse.WarehouseAttributes)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, warehouse, result)
		warehouseRepo.AssertExpectations(t)
		localityRepo.AssertExpectations(t)
	})

	t.Run("case 2: not found - warehouse not found", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}

		errNotFound := eh.GetErrNotFound(eh.WAREHOUSE)
		warehouseRepo.On("GetByID", warehouse.ID).Return(model.Warehouse{}, errNotFound)

		// Act
		result, err := warehouseService.Update(warehouse.ID, warehouse.WarehouseAttributes)

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
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}
		newWarehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesB,
		}

		errConflict := eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE)
		warehouseRepo.On("GetByID", warehouse.ID).Return(warehouse, nil)
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

	t.Run("case 4: foreign key - locality id doesn't exists", func(t *testing.T) {
		// Arrange
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		localityRepo := localityRepository.NewLocalityRepositoryMock()
		warehouseService := service.NewWarehouseService(warehouseRepo, localityRepo)
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}

		errForeignKey := eh.GetErrForeignKey(eh.LOCALITY)
		warehouseRepo.On("GetByID", warehouse.ID).Return(warehouse, nil)
		localityRepo.On("GetByID", *warehouse.LocalityID).Return(model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY))

		// Act
		result, err := warehouseService.Update(warehouse.ID, warehouse.WarehouseAttributes)

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
		warehouse := model.Warehouse{
			ID:                  1,
			WarehouseAttributes: warehouseAttributesA,
		}

		warehouseRepo.On("GetByID", warehouse.ID).Return(warehouse, nil)
		warehouseRepo.On("Delete", warehouse.ID).Return(nil)

		// Act
		err := warehouseService.Delete(warehouse.ID)

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
