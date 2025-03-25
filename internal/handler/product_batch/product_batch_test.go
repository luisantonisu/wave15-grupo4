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
	handler "github.com/luisantonisu/wave15-grupo4/internal/handler/product_batch"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/product_batch"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProductBatchHandler_Create(t *testing.T) {
	// Caso base válido
	validRequest := dto.ProductBatchRequestDTO{
		BatchNumber:        "B001",
		CurrentQuantity:    100,
		CurrentTemperature: 10.5,
		DueDate:            "2025-06-01",
		InitialQuantity:    100,
		ManufacturingDate:  "2025-05-01",
		ManufacturingHour:  "14:00",
		MinimumTemperature: 5.0,
		ProductID:          1,
		SectionID:          10,
	}

	validModel := model.ProductBatch{
		ID:                     1,
		ProductBatchAttributes: helper.ProductBatchRequestDTOToProductBatch(validRequest),
	}

	t.Run("caso 1: creación exitosa", func(t *testing.T) {
		mockService := service.NewMockProductBatchService()
		testHandler := handler.NewProductBatchHandler(mockService)

		mockService.On("Create", mock.Anything).Return(validModel, nil)

		body, _ := json.Marshal(validRequest)
		req := httptest.NewRequest(http.MethodPost, "/productBatches", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Post("/productBatches", testHandler.Create())
		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusCreated, rr.Code)
		require.Contains(t, rr.Body.String(), `"ID":1`)
		mockService.AssertExpectations(t)
	})

	t.Run("caso 2: cuerpo inválido (JSON malformado)", func(t *testing.T) {
		mockService := service.NewMockProductBatchService()
		testHandler := handler.NewProductBatchHandler(mockService)

		body := []byte(`{ invalid json `)
		req := httptest.NewRequest(http.MethodPost, "/productBatches", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Post("/productBatches", testHandler.Create())
		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
		require.Contains(t, rr.Body.String(), "invalid request body")
	})

	t.Run("caso 3: error del servicio (clave foránea no encontrada)", func(t *testing.T) {
		mockService := service.NewMockProductBatchService()
		testHandler := handler.NewProductBatchHandler(mockService)

		mockService.On("Create", mock.Anything).Return(model.ProductBatch{}, eh.GetErrForeignKey(eh.PRODUCT))

		body, _ := json.Marshal(validRequest)
		req := httptest.NewRequest(http.MethodPost, "/productBatches", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Post("/productBatches", testHandler.Create())
		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusConflict, rr.Code)
		require.Contains(t, rr.Body.String(), "product foreign key not found")
	})
}
