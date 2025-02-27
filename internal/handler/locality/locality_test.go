package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	handler "github.com/luisantonisu/wave15-grupo4/internal/handler/locality"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/locality"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLocalityHandler_CarriersReport(t *testing.T) {
	record1 := model.CarriersReport{
		LocalityID:    1,
		LocalityName:  "Locality 1",
		CarriersCount: 10,
	}
	record2 := model.CarriersReport{
		LocalityID:    2,
		LocalityName:  "Locality 2",
		CarriersCount: 1,
	}
	report := []model.CarriersReport{record1, record2}

	t.Run("case 1: get carriers report successfully", func(t *testing.T) {
		// Arrange
		localityService := service.NewLocalityServiceMock()
		localityHandler := handler.NewLocalityHandler(localityService)
		localityService.On("CarriersReport", mock.Anything).Return(report, nil)

		rt := chi.NewRouter()
		rt.Get("/localities/reportCarriers", localityHandler.CarriersReport())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/localities/reportCarriers", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[
			{"locality_id":1,"locality_name":"Locality 1","carriers_count":10},
			{"locality_id":2,"locality_name":"Locality 2","carriers_count":1}
		]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		localityService.AssertExpectations(t)
	})

	t.Run("case 2: get carriers report by locality id", func(t *testing.T) {
		// Arrange
		localityService := service.NewLocalityServiceMock()
		localityHandler := handler.NewLocalityHandler(localityService)
		localityService.On("CarriersReport", mock.Anything).Return([]model.CarriersReport{record1}, nil)

		rt := chi.NewRouter()
		rt.Get("/localities/reportCarriers", localityHandler.CarriersReport())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/localities/reportCarriers?id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[
			{"locality_id":1,"locality_name":"Locality 1","carriers_count":10}
		]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		localityService.AssertExpectations(t)
	})

	t.Run("case 3: not found - locality id not found", func(t *testing.T) {
		// Arrange
		localityService := service.NewLocalityServiceMock()
		localityHandler := handler.NewLocalityHandler(localityService)
		localityService.On("CarriersReport", mock.Anything).Return([]model.CarriersReport{}, eh.GetErrNotFound(eh.LOCALITY))

		rt := chi.NewRouter()
		rt.Get("/localities/reportCarriers", localityHandler.CarriersReport())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/localities/reportCarriers?id=3", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Not Found","message": "locality not found"}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		localityService.AssertExpectations(t)
	})

	t.Run("case 4: bad request - invalid locality id", func(t *testing.T) {
		// Arrange
		localityService := service.NewLocalityServiceMock()
		localityHandler := handler.NewLocalityHandler(localityService)

		rt := chi.NewRouter()
		rt.Get("/localities/reportCarriers", localityHandler.CarriersReport())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/localities/reportCarriers?id=invalid", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		localityService.AssertExpectations(t)
	})
}
