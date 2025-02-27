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
	handler "github.com/luisantonisu/wave15-grupo4/internal/handler/warehouse"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	warehouseCodeA  = "WH1"
	warehouseCodeB  = "WH2"
	address         = "Address 1"
	telephone       = uint(123456789)
	capacity        = 100
	temperature     = float32(20.5)
	localityID      = 1
	warehouseModelA = model.Warehouse{
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
	warehouseModelB = model.Warehouse{
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
	warehouseRequest = dto.WarehouseRequestDTO{
		WarehouseCode:      &warehouseCodeB,
		Address:            &address,
		Telephone:          &telephone,
		MinimumCapacity:    &capacity,
		MinimumTemperature: &temperature,
		LocalityID:         &localityID,
	}
)

func TestWarehouseHandler_GetAll(t *testing.T) {
	t.Run("case 1: get all warehouses successfully", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouses := []model.Warehouse{warehouseModelA, warehouseModelB}
		warehouseService.On("GetAll", mock.Anything).Return(warehouses, nil)

		rt := chi.NewRouter()
		rt.Get("/warehouses", warehouseHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouses", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[
			{"id":1,"warehouse_code":"WH1","address":"Address 1","telephone":123456789,"minimum_capacity":100,"minimum_temperature":20.5,"locality_id":1},
			{"id":2,"warehouse_code":"WH2","address":"Address 1","telephone":123456789,"minimum_capacity":100,"minimum_temperature":20.5,"locality_id":1}
		]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 2: no warehouses found", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("GetAll", mock.Anything).Return([]model.Warehouse{}, nil)

		rt := chi.NewRouter()
		rt.Get("/warehouses", warehouseHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouses", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 3: internal server error - error getting warehouses", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("GetAll", mock.Anything).Return([]model.Warehouse{}, eh.GetErrDatabase(eh.WAREHOUSE))

		rt := chi.NewRouter()
		rt.Get("/warehouses", warehouseHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouses", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Internal Server Error","message": "database error: warehouse"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})
}

func TestWarehouseHandler_GetByID(t *testing.T) {
	t.Run("case 1: get warehouse by id successfully", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("GetByID", 1).Return(warehouseModelA, nil)

		rt := chi.NewRouter()
		rt.Get("/warehouses/{id}", warehouseHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouses/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"warehouse_code":"WH1","address":"Address 1","telephone":123456789,"minimum_capacity":100,"minimum_temperature":20.5,"locality_id":1}}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 2: not found - get warehouse by id not found", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("GetByID", 1).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE))

		rt := chi.NewRouter()
		rt.Get("/warehouses/{id}", warehouseHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouses/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "warehouse not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 3: bad request - invalid id", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)

		rt := chi.NewRouter()
		rt.Get("/warehouses/{id}", warehouseHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/warehouses/abc", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})
}

func TestWarehouseHandler_Create(t *testing.T) {
	t.Run("case 1: create warehouse successfully", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Create", mock.Anything).Return(warehouseModelA, nil)

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/warehouses", warehouseHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"warehouse_code":"WH1","address":"Address 1","telephone":123456789,"minimum_capacity":100,"minimum_temperature":20.5,"locality_id":1}}`
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 2: bad request - invalid body", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)

		rt := chi.NewRouter()
		rt.Post("/warehouses", warehouseHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader([]byte("invalid body"))), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid request body"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 3: internal server error - error creating warehouse", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Create", mock.Anything).Return(model.Warehouse{}, eh.GetErrDatabase(eh.WAREHOUSE))

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/warehouses", warehouseHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Internal Server Error","message": "database error: warehouse"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 4: bad request - invalid locality id", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Create", mock.Anything).Return(model.Warehouse{}, eh.GetErrForeignKey(eh.LOCALITY))

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/warehouses", warehouseHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Conflict","message": "locality foreign key not found"}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 5: bad request - invalid warehouse code", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Create", mock.Anything).Return(model.Warehouse{}, eh.GetErrAlreadyExists(eh.WAREHOUSE_CODE))

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/warehouses", warehouseHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Conflict","message": "warehouse code already exists"}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 6: bad request - missing warehouse code", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Create", mock.Anything).Return(model.Warehouse{}, eh.GetErrInvalidData(eh.WAREHOUSE_CODE))

		warehouseReq := dto.WarehouseRequestDTO{
			WarehouseCode: nil,
		}
		body, err := json.Marshal(warehouseReq)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/warehouses", warehouseHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Unprocessable Entity","message": "invalid data: warehouse code"}`
		require.Equal(t, http.StatusUnprocessableEntity, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})
}

func TestWarehouseHandler_Update(t *testing.T) {
	t.Run("case 1: update warehouse successfully", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Update", 1, mock.Anything).Return(warehouseModelA, nil)

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Put("/warehouses/{id}", warehouseHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPut, "/warehouses/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"warehouse_code":"WH1","address":"Address 1","telephone":123456789,"minimum_capacity":100,"minimum_temperature":20.5,"locality_id":1}}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 2: bad request - invalid id", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Put("/warehouses/{id}", warehouseHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPut, "/warehouses/abc", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 3: bad request - invalid body", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)

		rt := chi.NewRouter()
		rt.Put("/warehouses/{id}", warehouseHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPut, "/warehouses/1", bytes.NewReader([]byte("invalid body"))), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid request body"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 4: internal server error - error updating warehouse", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Update", 1, mock.Anything).Return(model.Warehouse{}, eh.GetErrDatabase(eh.WAREHOUSE))

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Put("/warehouses/{id}", warehouseHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPut, "/warehouses/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Internal Server Error","message": "database error: warehouse"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 5: bad request - invalid locality id", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Update", 1, mock.Anything).Return(model.Warehouse{}, eh.GetErrForeignKey(eh.LOCALITY))

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Put("/warehouses/{id}", warehouseHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPut, "/warehouses/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Conflict","message": "locality foreign key not found"}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 6: not found - warehouse not found", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Update", 1, mock.Anything).Return(model.Warehouse{}, eh.GetErrNotFound(eh.WAREHOUSE))

		body, err := json.Marshal(warehouseRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Put("/warehouses/{id}", warehouseHandler.Update())

		// Act
		req, res := httptest.NewRequest(http.MethodPut, "/warehouses/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "warehouse not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})
}

func TestWarehouseHandler_Delete(t *testing.T) {
	t.Run("case 1: delete warehouse successfully", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Delete", 1).Return(nil)

		rt := chi.NewRouter()
		rt.Delete("/warehouses/{id}", warehouseHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/warehouses/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNoContent, res.Code)
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 2: not found - warehouse not found", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)
		warehouseService.On("Delete", 1).Return(eh.GetErrNotFound(eh.WAREHOUSE))

		rt := chi.NewRouter()
		rt.Delete("/warehouses/{id}", warehouseHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/warehouses/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "warehouse not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})

	t.Run("case 3: bad request - invalid id", func(t *testing.T) {
		// Arrange
		warehouseService := service.NewWarehouseServiceMock()
		warehouseHandler := handler.NewWarehouseHandler(warehouseService)

		rt := chi.NewRouter()
		rt.Delete("/warehouses/{id}", warehouseHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/warehouses/abc", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		warehouseService.AssertExpectations(t)
	})
}