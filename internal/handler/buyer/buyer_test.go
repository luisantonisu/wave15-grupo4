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
	handler "github.com/luisantonisu/wave15-grupo4/internal/handler/buyer"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	cardNumberId    = "123456"
	firstName       = "John"
	lastName        = "Doe"
	cardNumberIdB   = "654321"
	firstNameB      = "Jane"
	lastNameB       = "Doe"
	buyerAttributes = model.BuyerAttributes{
		CardNumberId: &cardNumberId,
		FirstName:    &firstName,
		LastName:     &lastName,
	}
	buyer = model.Buyer{
		ID:              1,
		BuyerAttributes: buyerAttributes,
	}
	buyerAttributesB = model.BuyerAttributes{
		CardNumberId: &cardNumberIdB,
		FirstName:    &firstNameB,
		LastName:     &lastNameB,
	}
	buyerB = model.Buyer{
		ID:              2,
		BuyerAttributes: buyerAttributesB,
	}
	buyerRequest = dto.BuyerRequestDTO{
		CardNumberId: &cardNumberId,
		FirstName:    &firstName,
		LastName:     &lastName,
	}
	buyerRequestUpdate = dto.BuyerRequestDTO{
		FirstName: &firstName,
		LastName:  &lastName,
	}
	buyerUpdate = model.Buyer{
		ID: 1,
		BuyerAttributes: model.BuyerAttributes{
			CardNumberId: &cardNumberId,
			FirstName:    &firstNameB,
			LastName:     &lastNameB,
		},
	}
	PurchaseOrderReport = model.ReportPurchaseOrders{
		ID:                  1,
		CardNumberId:        cardNumberId,
		FirstName:           firstName,
		LastName:            lastName,
		PurchaseOrdersCount: 1,
	}
	PurchaseOrderReportB = model.ReportPurchaseOrders{
		ID:                  2,
		CardNumberId:        cardNumberIdB,
		FirstName:           firstNameB,
		LastName:            lastNameB,
		PurchaseOrdersCount: 2,
	}
)

func TestBuyerHandler_Create(t *testing.T) {
	t.Run("case 1: create buyer successfully", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("Create", mock.Anything).Return(buyer, nil)

		body, err := json.Marshal(buyerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/buyers", buyerHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/buyers", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe"}}`
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
	t.Run("case 2: unprocesable entity - invalid card number id", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)

		body, err := json.Marshal(dto.BuyerRequestDTO{
			CardNumberId: nil,
			FirstName:    &firstName,
			LastName:     &lastName,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/buyers", buyerHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/buyers", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Unprocessable Entity","message":"invalid data: card number ID"}`
		require.Equal(t, http.StatusUnprocessableEntity, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertNotCalled(t, "Create")
	})
	t.Run("case 3: bad request - invalid request body", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)

		body, err := json.Marshal(`{
			CardNumberId: 1,
			FirstName:    &firstName,
			LastName:     &lastName,
			}`)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/buyers", buyerHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/buyers", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message":"invalid request body"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertNotCalled(t, "Create")
	})
	t.Run("case 4: internal server error - service error", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("Create", mock.Anything).Return(model.Buyer{}, eh.GetErrInternalServer(eh.BUYER))

		body, err := json.Marshal(buyerRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/buyers", buyerHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/buyers", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Internal Server Error","message":"internal server error"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
}

func TestBuyerHandler_GetAll(t *testing.T) {
	t.Run("case 1: get all buyers successfully", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyers := []model.Buyer{buyer, buyerB}
		buyerService.On("GetAll").Return(buyers, nil)

		rt := chi.NewRouter()
		rt.Get("/buyers", buyerHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[
			{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe"},
			{"id":2,"card_number_id":"654321","first_name":"Jane","last_name":"Doe"}
		]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})

	t.Run("case 2: get all buyers successfully, no buyers found", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("GetAll").Return([]model.Buyer{}, nil)

		rt := chi.NewRouter()
		rt.Get("/buyers", buyerHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data": null}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
	t.Run("case 3: internal Server Error - service error", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("GetAll").Return([]model.Buyer{}, eh.GetErrInternalServer(eh.BUYER))

		rt := chi.NewRouter()
		rt.Get("/buyers", buyerHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Internal Server Error","message": "internal server error"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
}

func TestBuyerHandler_GetByID(t *testing.T) {
	t.Run("case 1: get buyer by id successfully", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("GetByID", 1).Return(buyer, nil)

		rt := chi.NewRouter()
		rt.Get("/buyers/{id}", buyerHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe"}}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
	t.Run("case 2: not found - get buyer by id, id non existent", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("GetByID", 1).Return(model.Buyer{}, eh.GetErrNotFound(eh.BUYER))

		rt := chi.NewRouter()
		rt.Get("/buyers/{id}", buyerHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "buyer not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
	t.Run("case 3: bad request - get buyer by, invalid id", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)

		rt := chi.NewRouter()
		rt.Get("/buyers/{id}", buyerHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers/invalidId", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertNotCalled(t, "GetByID")
	})
}

func TestBuyerHandler_Update(t *testing.T) {
	t.Run("case 1: update buyer successfully", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("Update", 1, mock.Anything).Return(buyerUpdate, nil)

		body, err := json.Marshal(buyerRequestUpdate)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/buyers/{id}", buyerHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/buyers/1", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"card_number_id":"123456","first_name":"Jane","last_name":"Doe"}}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
	t.Run("case 2: not found - update buyer, id non existent", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("Update", 3, mock.Anything).Return(model.Buyer{}, eh.GetErrNotFound(eh.BUYER))

		body, err := json.Marshal(buyerRequestUpdate)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/buyers/{id}", buyerHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/buyers/3", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "buyer not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
	t.Run("case 3: bad request - update buyer, invalid id", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("Update", "invalidId", mock.Anything).Return(model.Buyer{}, eh.INVALID_ID)

		body, err := json.Marshal(buyerRequestUpdate)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/buyers/{id}", buyerHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/buyers/invalidId", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertNotCalled(t, "Update")
	})

	t.Run("case 4: bad request - update buyer, invalid body", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("Update", 1, mock.Anything).Return(model.Buyer{}, eh.INVALID_BODY)

		body, err := json.Marshal(`{
			CardNumberId: 1,
			FirstName:    &firstName,
			LastName:     &lastName,
		}`)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/buyers/{id}", buyerHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/buyers/1", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message":"invalid request body"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertNotCalled(t, "Update")
	})
}

func TestBuyerHandler_Delete(t *testing.T) {
	t.Run("case 1: delete buyer successfully", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("Delete", 1).Return(nil)

		rt := chi.NewRouter()
		rt.Delete("/buyers/{id}", buyerHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/buyers/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNoContent, res.Code)
		buyerService.AssertExpectations(t)
	})
	t.Run("case 2: not found - delete buyer, id non existent", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		buyerService.On("Delete", 3).Return(eh.GetErrNotFound(eh.BUYER))

		rt := chi.NewRouter()
		rt.Delete("/buyers/{id}", buyerHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/buyers/3", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "buyer not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})

	t.Run("case 3: bad request - delete buyer, invalid id", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)

		rt := chi.NewRouter()
		rt.Delete("/buyers/{id}", buyerHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/buyers/invalidId", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertNotCalled(t, "Delete")
	})
}

func TestBuyerHanlder_Report(t *testing.T) {
	t.Run("case 1: get purchase order report successfully, all buyers", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		report := []model.ReportPurchaseOrders{PurchaseOrderReport, PurchaseOrderReportB}
		var id *int
		buyerService.On("PurchaseOrderReport", id).Return(report, nil)

		rt := chi.NewRouter()
		rt.Get("/buyers/reportPurchaseOrders", buyerHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers/reportPurchaseOrders", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[
			{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe","purchase_orders_count":1},
			{"id":2,"card_number_id":"654321","first_name":"Jane","last_name":"Doe","purchase_orders_count":2}
		]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
	t.Run("case 2: get purchase order report successfully, specific id", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		report := []model.ReportPurchaseOrders{PurchaseOrderReport}
		id := 1
		buyerService.On("PurchaseOrderReport", &id).Return(report, nil)

		rt := chi.NewRouter()
		rt.Get("/buyers/reportPurchaseOrders", buyerHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers/reportPurchaseOrders?id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[
			{"id":1,"card_number_id":"123456","first_name":"John","last_name":"Doe","purchase_orders_count":1}
		]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})

	t.Run("case 3: get purchase order report successfully but empty, specific id", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		id := 1
		buyerService.On("PurchaseOrderReport", &id).Return([]model.ReportPurchaseOrders{}, nil)

		rt := chi.NewRouter()
		rt.Get("/buyers/reportPurchaseOrders", buyerHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers/reportPurchaseOrders?id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
	t.Run("case 4: bad request -  get purchase order report, invalid id", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)

		rt := chi.NewRouter()
		rt.Get("/buyers/reportPurchaseOrders", buyerHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers/reportPurchaseOrders?id=invalidId", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertNotCalled(t, "PurchaseOrderReport")
	})
	t.Run("case 5: not found - get purchase order report, id non existent", func(t *testing.T) {
		// Arrange
		buyerService := service.NewBuyerServiceMock()
		buyerHandler := handler.NewBuyerHandler(buyerService)
		id := 1
		buyerService.On("PurchaseOrderReport", &id).Return([]model.ReportPurchaseOrders{}, eh.GetErrNotFound(eh.BUYER))

		rt := chi.NewRouter()
		rt.Get("/buyers/reportPurchaseOrders", buyerHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/buyers/reportPurchaseOrders?id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "buyer not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		buyerService.AssertExpectations(t)
	})
}
