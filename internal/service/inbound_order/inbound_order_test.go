package service

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repositoryEm "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
	repositoryInb "github.com/luisantonisu/wave15-grupo4/internal/repository/inbound_order"
	repositoryWh "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	date              = "2025-08-10 09:40:18"
	orderNumberA      = 12345
	employeeID        = 1
	warehouseID       = 1
	productBatchID    = 1
	inboundOrderMockA = model.InboundOrder{
		ID: 1,
		InboundOrderAttributes: model.InboundOrderAttributes{
			OrderDate:      &date,
			OrderNumber:    &orderNumberA,
			EmployeeID:     &employeeID,
			ProductBatchID: &productBatchID,
			WarehouseID:    &warehouseID,
		},
	}
	inboundOrderMockB = model.InboundOrderAttributes{
		OrderDate:      nil,
		OrderNumber:    &orderNumberA,
		EmployeeID:     nil,
		ProductBatchID: nil,
		WarehouseID:    nil,
	}
)

func TestInboundOrderService_Create(t *testing.T) {

	t.Run("case 1: create inbound order successfully", func(t *testing.T) {
		//Arrange
		mockInboundOrderRepo := repositoryInb.NewInboundOrderRepoMock()
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewInboundOrderService(mockInboundOrderRepo, mockEmployeeRepo, mockWarehouseRepo)

		mockEmployeeRepo.On("GetByID", mock.Anything).Return(model.Employee{ID: 1}, nil)
		mockWarehouseRepo.On("GetByID", mock.Anything).Return(model.Warehouse{ID: 1}, nil)
		mockInboundOrderRepo.On("AlreadyExists", "product_batch_id", productBatchID).Return(true)
		mockInboundOrderRepo.On("AlreadyExists", "order_number", orderNumberA).Return(false)
		
		mockInboundOrderRepo.On("CreateInboundOrder", mock.Anything).Return(inboundOrderMockA, nil)

		//Act
		inboundOrder, err := service.Create(inboundOrderMockA.InboundOrderAttributes)

		//Assert
		require.NoError(t, err)
		require.Equal(t, inboundOrderMockA, inboundOrder)
		mockInboundOrderRepo.AssertExpectations(t)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 2: invalid data - employee missing fields", func(t *testing.T) {
		//Arrange
		mockInboundOrderRepo := repositoryInb.NewInboundOrderRepoMock()
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewInboundOrderService(mockInboundOrderRepo, mockEmployeeRepo, mockWarehouseRepo)

		//Act
		inboundOrder, err := service.Create(inboundOrderMockB)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, eh.GetErrInvalidData(eh.INBOUND_ORDER), err)
		require.Equal(t, model.InboundOrder{}, inboundOrder)
		mockInboundOrderRepo.AssertExpectations(t)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 3: foreign key - warehouse id not found", func(t *testing.T) {
		//Arrange
		mockInboundOrderRepo := repositoryInb.NewInboundOrderRepoMock()
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewInboundOrderService(mockInboundOrderRepo, mockEmployeeRepo, mockWarehouseRepo)

		mockEmployeeRepo.On("GetByID", mock.Anything).Return(model.Employee{ID: 1}, nil)
		mockWarehouseRepo.On("GetByID", mock.Anything).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE))

		//Act
		inboundOrder, err := service.Create(inboundOrderMockA.InboundOrderAttributes)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.WAREHOUSE), err)
		require.Equal(t, model.InboundOrder{}, inboundOrder)
		mockInboundOrderRepo.AssertExpectations(t)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 4: foreign key - employee id not found", func(t *testing.T) {
		//Arrange
		mockInboundOrderRepo := repositoryInb.NewInboundOrderRepoMock()
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewInboundOrderService(mockInboundOrderRepo, mockEmployeeRepo, mockWarehouseRepo)

		mockEmployeeRepo.On("GetByID", mock.Anything).Return(model.Employee{}, eh.GetErrForeignKey(eh.EMPLOYEE))

		//Act
		inboundOrder, err := service.Create(inboundOrderMockA.InboundOrderAttributes)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.EMPLOYEE), err)
		require.Equal(t, model.InboundOrder{}, inboundOrder)
		mockInboundOrderRepo.AssertExpectations(t)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 5: foreign key - product batch id not found", func(t *testing.T) {
		//Arrange
		mockInboundOrderRepo := repositoryInb.NewInboundOrderRepoMock()
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewInboundOrderService(mockInboundOrderRepo, mockEmployeeRepo, mockWarehouseRepo)

		mockEmployeeRepo.On("GetByID", mock.Anything).Return(model.Employee{ID: 1}, nil)
		mockWarehouseRepo.On("GetByID", mock.Anything).Return(model.Warehouse{ID: 1}, nil)
		mockInboundOrderRepo.On("AlreadyExists", "product_batch_id", productBatchID).Return(false)

		//Act
		inboundOrder, err := service.Create(inboundOrderMockA.InboundOrderAttributes)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.PRODUCT_BATCH_ID), err)
		require.Equal(t, model.InboundOrder{}, inboundOrder)
		mockInboundOrderRepo.AssertExpectations(t)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

	t.Run("case 6: conflict - order number duplicated", func(t *testing.T) {
		//Arrange
		mockInboundOrderRepo := repositoryInb.NewInboundOrderRepoMock()
		mockEmployeeRepo := repositoryEm.NewEmployeeMockRepository()
		mockWarehouseRepo := repositoryWh.NewWarehouseRepositoryMock()
		service := NewInboundOrderService(mockInboundOrderRepo, mockEmployeeRepo, mockWarehouseRepo)

		mockEmployeeRepo.On("GetByID", mock.Anything).Return(model.Employee{ID: 1}, nil)
		mockWarehouseRepo.On("GetByID", mock.Anything).Return(model.Warehouse{ID: 1}, nil)
		mockInboundOrderRepo.On("AlreadyExists", "product_batch_id", productBatchID).Return(true)
		mockInboundOrderRepo.On("AlreadyExists", "order_number", orderNumberA).Return(true)
		
		//Act
		inboundOrder, err := service.Create(inboundOrderMockA.InboundOrderAttributes)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, eh.GetErrAlreadyExistsCompose(eh.INBOUND_ORDER, eh.ORDER_NUMBER), err)
		require.Equal(t, model.InboundOrder{}, inboundOrder)
		mockInboundOrderRepo.AssertExpectations(t)
		mockEmployeeRepo.AssertExpectations(t)
		mockWarehouseRepo.AssertExpectations(t)
	})

}
