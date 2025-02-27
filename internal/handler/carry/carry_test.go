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
	handler "github.com/luisantonisu/wave15-grupo4/internal/handler/carry"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/carry"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	carryID     = "C1"
	companyName = "Company 1"
	address     = "Address 1"
	telephone   = uint(123456789)
	LocalityID  = 1

	carryModel = model.Carry{
		ID: 1,
		CarryAttributes: model.CarryAttributes{
			CarryID:     &carryID,
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
			LocalityID:  &LocalityID,
		},
	}

	carryRequest = dto.CarryRequestDTO{
		CarryID:     &carryID,
		CompanyName: &companyName,
		Address:     &address,
		Telephone:   &telephone,
		LocalityID:  &LocalityID,
	}
)

func TestCarryHandler_Create(t *testing.T) {
	t.Run("case 1: create a carry successfully", func(t *testing.T) {
		// Arrange
		carryService := service.NewCarryServiceMock()
		carryHandler := handler.NewCarryHandler(carryService)
		carryService.On("Create", mock.Anything).Return(carryModel, nil)

		body, err := json.Marshal(carryRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/carriers", carryHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/carriers", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":{"id":1,"carry_id":"C1","company_name":"Company 1","address":"Address 1","telephone":123456789,"locality_id":1}}`
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		carryService.AssertExpectations(t)
	})

	t.Run("case 2: bad request - invalid body", func(t *testing.T) {
		// Arrange
		carryService := service.NewCarryServiceMock()
		carryHandler := handler.NewCarryHandler(carryService)

		rt := chi.NewRouter()
		rt.Post("/carriers", carryHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/carriers", bytes.NewReader([]byte("invalid body"))), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid request body"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		carryService.AssertExpectations(t)
	})

	t.Run("case 3: internal server error - service error", func(t *testing.T) {
		// Arrange
		carryService := service.NewCarryServiceMock()
		carryHandler := handler.NewCarryHandler(carryService)
		carryService.On("Create", mock.Anything).Return(model.Carry{}, eh.GetErrDatabase(eh.CARRY))

		body, err := json.Marshal(carryRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/carriers", carryHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/carriers", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Internal Server Error","message": "database error: carry"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		carryService.AssertExpectations(t)
	})
}
