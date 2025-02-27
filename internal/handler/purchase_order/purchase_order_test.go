package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	handler "github.com/luisantonisu/wave15-grupo4/internal/handler/purchase_order"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/purchase_order"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	orderNumber   = "123456"
	orderDate     = "2021-09-01"
	trackingCode  = "ABC123"
	buyerID       = 1
	carrierID     = 1
	orderStatusID = 1
	warehouseID   = 1

	purchaseOrderAttributes = model.PurchaseOrderAttributes{
		OrderNumber:   &orderNumber,
		OrderDate:     &orderDate,
		TrackingCode:  &trackingCode,
		BuyerID:       &buyerID,
		CarrierID:     &carrierID,
		OrderStatusID: &orderStatusID,
		WarehouseID:   &warehouseID,
	}

	purchaseOrderRequestDTO = dto.PurchaseOrderRequestDTO{
		OrderNumber:   &orderNumber,
		OrderDate:     &orderDate,
		TrackingCode:  &trackingCode,
		BuyerID:       &buyerID,
		CarrierID:     &carrierID,
		OrderStatusID: &orderStatusID,
		WarehouseID:   &warehouseID,
	}

	purchaseOrder = model.PurchaseOrder{
		ID:                      1,
		PurchaseOrderAttributes: purchaseOrderAttributes,
	}
)

func TestCreatePurchaseOrder(t *testing.T) {
	t.Run("case 1: create purchase order successfully", func(t *testing.T) {
		// Arrange
		purchaseOrderService := service.NewPurchaseOrderServiceMock()
		purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)
		purchaseOrderService.On("Create", mock.Anything).Return(purchaseOrder, nil)

		body, err := json.Marshal(purchaseOrderRequestDTO)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/purchaseOrders", purchaseOrderHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/purchaseOrders", bytes.NewBuffer(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"order_number":"123456","order_date":"2021-09-01","tracking_code":"ABC123","buyer_id":1,"carrier_id":1,"order_status_id":1,"warehouse_id":1}}`
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		purchaseOrderService.AssertExpectations(t)
	})
	t.Run("case 2: unprocesable entity - invalid order number", func(t *testing.T) {
		// Arrange
		purchaseOrderService := service.NewPurchaseOrderServiceMock()
		purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

		body, err := json.Marshal(dto.PurchaseOrderRequestDTO{
			OrderNumber:   nil,
			OrderDate:     &orderDate,
			TrackingCode:  &trackingCode,
			BuyerID:       &buyerID,
			CarrierID:     &carrierID,
			OrderStatusID: &orderStatusID,
			WarehouseID:   &warehouseID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/purchaseOrders", purchaseOrderHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/purchaseOrders", bytes.NewBuffer(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Unprocessable Entity","message":"invalid data: order number"}`
		require.Equal(t, http.StatusUnprocessableEntity, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		purchaseOrderService.AssertExpectations(t)
	})
	t.Run("case 3: conflict - card number id already exists", func(t *testing.T) {
		// Arrange
		purchaseOrderService := service.NewPurchaseOrderServiceMock()
		purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)
		purchaseOrderService.On("Create", mock.Anything).Return(model.PurchaseOrder{}, eh.GetErrAlreadyExists(eh.ORDER_NUMBER))

		body, err := json.Marshal(purchaseOrderRequestDTO)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/purchaseOrders", purchaseOrderHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/purchaseOrders", bytes.NewBuffer(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Conflict","message":"order number already exists"}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		purchaseOrderService.AssertExpectations(t)
	})
	t.Run("case 4: bad request - invalid body", func(t *testing.T) {
		// Arrange
		purchaseOrderService := service.NewPurchaseOrderServiceMock()
		purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

		body, err := json.Marshal(`{
			OrderNumber: 123456,
			OrderDate: "2021-09-01",
			TrackingCode: "ABC123",
			BuyerID: 1,
			CarrierID: 1,
			OrderStatusID: 1,
			WarehouseID: 1,
		}`)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/purchaseOrders", purchaseOrderHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/purchaseOrders", bytes.NewBuffer(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message":"invalid request body"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		purchaseOrderService.AssertExpectations(t)
	})
}
