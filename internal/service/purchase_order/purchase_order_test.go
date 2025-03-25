package service_test

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	buyerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	carryRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/carry"
	orderStatusRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/order_status"
	purchaseOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/purchase_order"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/purchase_order"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

// Variables for testing
var (
	orderNumber   = "123456"
	orderDate     = "2021-09-01"
	trackingCode  = "TRACK123"
	buyerId       = 1
	carryId       = 1
	orderStatusId = 1
	warehouseId   = 1

	purchaseOrderAttributes = model.PurchaseOrderAttributes{
		OrderNumber:   &orderNumber,
		OrderDate:     &orderDate,
		TrackingCode:  &trackingCode,
		BuyerID:       &buyerId,
		CarrierID:     &carryId,
		OrderStatusID: &orderStatusId,
		WarehouseID:   &warehouseId,
	}
)

func TestPurchaseOrderService_Create(t *testing.T) {
	t.Run("case 1: create purchase order successfully", func(t *testing.T) {
		// Arrange
		purchaseOrderRepo := purchaseOrderRepository.NewPurchaseOrderRepositoryMock()
		buyerRepo := buyerRepository.NewBuyerRepositoryMock()
		carryRepo := carryRepository.NewCarryRepositoryMock()
		orderStatusRepo := orderStatusRepository.NewOrderStatusRepositoryMock()
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()
		purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo, buyerRepo, carryRepo, orderStatusRepo, warehouseRepo)
		purchaseOrder := model.PurchaseOrder{
			ID:                      1,
			PurchaseOrderAttributes: purchaseOrderAttributes,
		}
		purchaseOrderRepo.On("OrderNumberExists", "123456").Return(false)
		purchaseOrderRepo.On("Create", purchaseOrderAttributes).Return(purchaseOrder, nil)
		buyerRepo.On("GetByID", buyerId).Return(model.Buyer{}, nil)
		carryRepo.On("GetByID", carryId).Return(model.Carry{}, nil)
		orderStatusRepo.On("GetByID", orderStatusId).Return(model.OrderStatus{}, nil)
		warehouseRepo.On("GetByID", warehouseId).Return(model.Warehouse{}, nil)

		// Act
		purchaseOrderResult, err := purchaseOrderService.Create(purchaseOrderAttributes)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, purchaseOrderResult)
		require.Equal(t, purchaseOrder, purchaseOrderResult)
		purchaseOrderRepo.AssertExpectations(t)
		buyerRepo.AssertExpectations(t)
		carryRepo.AssertExpectations(t)
		orderStatusRepo.AssertExpectations(t)
		warehouseRepo.AssertExpectations(t)
	})
	t.Run("case 2: conflict - buyer foreing key not found", func(t *testing.T) {
		// Arrange
		purchaseOrderRepo := purchaseOrderRepository.NewPurchaseOrderRepositoryMock()
		buyerRepo := buyerRepository.NewBuyerRepositoryMock()
		carryRepo := carryRepository.NewCarryRepositoryMock()
		orderStatusRepo := orderStatusRepository.NewOrderStatusRepositoryMock()
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()

		purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo, buyerRepo, carryRepo, orderStatusRepo, warehouseRepo)
		purchaseOrderRepo.On("OrderNumberExists", "123456").Return(false)
		buyerRepo.On("GetByID", buyerId).Return(model.Buyer{}, eh.GetErrForeignKey(eh.BUYER))

		// Act
		purchaseOrderResult, err := purchaseOrderService.Create(purchaseOrderAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.BUYER), err)
		require.Equal(t, model.PurchaseOrder{}, purchaseOrderResult)
		purchaseOrderRepo.AssertExpectations(t)
		buyerRepo.AssertExpectations(t)
		carryRepo.AssertExpectations(t)
		orderStatusRepo.AssertExpectations(t)
		warehouseRepo.AssertExpectations(t)
	})
	t.Run("case 3: conflict - carry foreing key not found", func(t *testing.T) {
		// Arrange
		purchaseOrderRepo := purchaseOrderRepository.NewPurchaseOrderRepositoryMock()
		buyerRepo := buyerRepository.NewBuyerRepositoryMock()
		orderStatusRepo := orderStatusRepository.NewOrderStatusRepositoryMock()
		carryRepo := carryRepository.NewCarryRepositoryMock()
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()

		purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo, buyerRepo, carryRepo, orderStatusRepo, warehouseRepo)
		purchaseOrderRepo.On("OrderNumberExists", "123456").Return(false)
		buyerRepo.On("GetByID", buyerId).Return(model.Buyer{}, nil)
		carryRepo.On("GetByID", carryId).Return(model.Carry{}, eh.GetErrForeignKey(eh.CARRY))

		// Act
		purchaseOrderResult, err := purchaseOrderService.Create(purchaseOrderAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.CARRY), err)
		require.Equal(t, model.PurchaseOrder{}, purchaseOrderResult)
		purchaseOrderRepo.AssertExpectations(t)
		buyerRepo.AssertExpectations(t)
		carryRepo.AssertExpectations(t)
		orderStatusRepo.AssertExpectations(t)
		warehouseRepo.AssertExpectations(t)
	})
	t.Run("case 4: conflict - order status foreing key not found", func(t *testing.T) {
		// Arrange
		purchaseOrderRepo := purchaseOrderRepository.NewPurchaseOrderRepositoryMock()
		buyerRepo := buyerRepository.NewBuyerRepositoryMock()
		carryRepo := carryRepository.NewCarryRepositoryMock()
		orderStatusRepo := orderStatusRepository.NewOrderStatusRepositoryMock()
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()

		purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo, buyerRepo, carryRepo, orderStatusRepo, warehouseRepo)
		purchaseOrderRepo.On("OrderNumberExists", "123456").Return(false)
		buyerRepo.On("GetByID", buyerId).Return(model.Buyer{}, nil)
		carryRepo.On("GetByID", carryId).Return(model.Carry{}, nil)
		orderStatusRepo.On("GetByID", orderStatusId).Return(model.OrderStatus{}, eh.GetErrForeignKey(eh.ORDER_STATUS))

		// Act
		purchaseOrderResult, err := purchaseOrderService.Create(purchaseOrderAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.ORDER_STATUS), err)
		require.Equal(t, model.PurchaseOrder{}, purchaseOrderResult)
		purchaseOrderRepo.AssertExpectations(t)
		buyerRepo.AssertExpectations(t)
		carryRepo.AssertExpectations(t)
		orderStatusRepo.AssertExpectations(t)
		warehouseRepo.AssertExpectations(t)
	})
	t.Run("case 5: conflict - warehouse foreing key not found", func(t *testing.T) {
		// Arrange
		purchaseOrderRepo := purchaseOrderRepository.NewPurchaseOrderRepositoryMock()
		buyerRepo := buyerRepository.NewBuyerRepositoryMock()
		carryRepo := carryRepository.NewCarryRepositoryMock()
		orderStatusRepo := orderStatusRepository.NewOrderStatusRepositoryMock()
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()

		purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo, buyerRepo, carryRepo, orderStatusRepo, warehouseRepo)
		purchaseOrderRepo.On("OrderNumberExists", "123456").Return(false)
		buyerRepo.On("GetByID", buyerId).Return(model.Buyer{}, nil)
		carryRepo.On("GetByID", carryId).Return(model.Carry{}, nil)
		orderStatusRepo.On("GetByID", orderStatusId).Return(model.OrderStatus{}, nil)
		warehouseRepo.On("GetByID", warehouseId).Return(model.Warehouse{}, eh.GetErrForeignKey(eh.WAREHOUSE))

		// Act
		purchaseOrderResult, err := purchaseOrderService.Create(purchaseOrderAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, eh.GetErrForeignKey(eh.WAREHOUSE), err)
		require.Equal(t, model.PurchaseOrder{}, purchaseOrderResult)
		purchaseOrderRepo.AssertExpectations(t)
		buyerRepo.AssertExpectations(t)
		carryRepo.AssertExpectations(t)
		orderStatusRepo.AssertExpectations(t)
		warehouseRepo.AssertExpectations(t)
	})
	t.Run("case 6: conflict - purchase order number already exists", func(t *testing.T) {
		// Arrange
		purchaseOrderRepo := purchaseOrderRepository.NewPurchaseOrderRepositoryMock()
		buyerRepo := buyerRepository.NewBuyerRepositoryMock()
		carryRepo := carryRepository.NewCarryRepositoryMock()
		orderStatusRepo := orderStatusRepository.NewOrderStatusRepositoryMock()
		warehouseRepo := warehouseRepository.NewWarehouseRepositoryMock()

		purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo, buyerRepo, carryRepo, orderStatusRepo, warehouseRepo)
		purchaseOrderRepo.On("OrderNumberExists", "123456").Return(true)

		// Act
		purchaseOrderResult, err := purchaseOrderService.Create(purchaseOrderAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, eh.GetErrAlreadyExists(eh.ORDER_NUMBER), err)
		require.Equal(t, model.PurchaseOrder{}, purchaseOrderResult)
		purchaseOrderRepo.AssertExpectations(t)
		buyerRepo.AssertExpectations(t)
		carryRepo.AssertExpectations(t)
		orderStatusRepo.AssertExpectations(t)
		warehouseRepo.AssertExpectations(t)
	})
}
