package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/inbound_order"
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
)

func TestInboundOrderHandler_Create(t *testing.T) {
	t.Run("case 1: create inbound order successfully", func(t *testing.T) {
		// Arrange
		inboundService := service.NewInboundOrderServiceMock()
		inboundHandler := NewInboundOrderHandler(inboundService)

		inboundService.On("Create", mock.Anything).Return(inboundOrderMockA, nil)

		body, err := json.Marshal(inboundOrderMockA.InboundOrderAttributes)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/inboundOrders", inboundHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/inboundOrders", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": {"id":1,"order_date":"2025-08-10 09:40:18","order_number":12345,"employee_id":1,"product_batch_id":1,"warehouse_id":1}
		}`

		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 2: invalid data - problems creating inbound order", func(t *testing.T) {
		// Arrange
		inboundService := service.NewInboundOrderServiceMock()
		inboundHandler := NewInboundOrderHandler(inboundService)

		inboundService.On("Create", mock.Anything).Return(model.InboundOrder{}, eh.GetErrInvalidData(eh.INBOUND_ORDER))

		body, err := json.Marshal(dto.InboundOrderRequestDTO{
			OrderDate:      nil,
			OrderNumber:    nil,
			EmployeeID:     nil,
			ProductBatchID: &productBatchID,
			WarehouseID:    &warehouseID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/inboundOrders", inboundHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/inboundOrders", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Unprocessable Entity",
			"message": "invalid data: inbound order"
		}`

		require.Equal(t, http.StatusUnprocessableEntity, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 3: bad request - invalid body", func(t *testing.T) {
		// Arrange
		inboundOrderSv := service.NewInboundOrderServiceMock()
		iboundOrderHnd := NewInboundOrderHandler(inboundOrderSv)
		inboundOrderSv.On("Create", mock.Anything).Return(model.InboundOrder{}, eh.INVALID_BODY)

		body, err := json.Marshal(`{
			"order_date": "2025-08-10 09:40:18",
			"order_number": 12345,
			"employee_id": 1,
			"product_batch_id": 1,
			"warehouse_id": 1
		}`)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/inboundOrders", iboundOrderHnd.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/inboundOrders", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Bad Request",
			"message": "invalid request body"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}
