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
	handler "github.com/luisantonisu/wave15-grupo4/internal/handler/locality"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/locality"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	localityID   = "1"
	localityName = "Merida"
	provinceName = "Yucatan"
	countryName  = "Mexico"

	report1 = model.LocalityReport{
		Id:           1,
		LocalityName: "Callao",
		SellerCount:  10,
	}
	report2 = model.LocalityReport{
		Id:           2,
		LocalityName: "Merida",
		SellerCount:  9,
	}

	SellerReport = []model.LocalityReport{report1, report2}

	localityRequest = dto.LocalityRequestDTO{
		Data: dto.LocalityDataDTO{
			Id:           &localityID,
			LocalityName: &localityName,
			ProvinceName: &provinceName,
			CountryName:  &countryName,
		},
	}

	locality1 = model.Locality{
		Id: localityID,
		LocalityAttributes: model.LocalityAttributes{
			LocalityName: &localityName,
			ProvinceName: &provinceName,
			CountryName:  &countryName,
		},
	}
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

func TestLocalityHandler_SellersReport(t *testing.T) {
	t.Run("get sellers report successfully", func(t *testing.T) {
		// Arrange
		localityService := service.NewLocalityServiceMock()
		localityHandler := handler.NewLocalityHandler(localityService)
		localityService.On("SellersReport", mock.Anything).Return(SellerReport, nil)

		rt := chi.NewRouter()
		rt.Get("/localities/reportSellers", localityHandler.SellersReport())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/localities/reportSellers", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[
			{"locality_id":"1","locality_name":"Callao","sellers_count":10},
			{"locality_id":"2","locality_name":"Merida","sellers_count":9}
		]}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		localityService.AssertExpectations(t)
	})
	t.Run("case 2: get sellers report by locality id", func(t *testing.T) {
		// Arrange
		localityService := service.NewLocalityServiceMock()
		localityHandler := handler.NewLocalityHandler(localityService)
		localityService.On("SellersReport", mock.Anything).Return([]model.LocalityReport{report1}, nil)

		rt := chi.NewRouter()
		rt.Get("/localities/reportSellers", localityHandler.SellersReport())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/localities/reportSellers?id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"data":[
			{"locality_id":"1","locality_name":"Callao","sellers_count":10}
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
		localityService.On("SellersReport", mock.Anything).Return([]model.LocalityReport{}, eh.GetErrNotFound(eh.LOCALITY))

		rt := chi.NewRouter()
		rt.Get("/localities/reportSellers", localityHandler.SellersReport())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/localities/reportSellers?id=3", nil), httptest.NewRecorder()
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
		rt.Get("/localities/reportSellers", localityHandler.SellersReport())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/localities/reportSellers?id=invalid", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Bad Request","message": "invalid id"}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		localityService.AssertExpectations(t)
	})
}

func TestLocalityHandler_Create(t *testing.T) {
	t.Run("case 1: locality create successfully", func(t *testing.T) {
		// Arrange
		localityService := service.NewLocalityServiceMock()
		localityHandler := handler.NewLocalityHandler(localityService)
		localityService.On("Create", mock.Anything).Return(locality1, nil)

		body, err := json.Marshal(localityRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/localities", localityHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/localities", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{ "data":{
        	"id": "1",
       	    "locality_name": "Merida",
            "province_name": "Yucatan",
            "country_name": "Mexico"
        }}`
		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		localityService.AssertExpectations(t)
	})
	t.Run("case 2: internal server error", func(t *testing.T) {
		// Arrange
		localityService := service.NewLocalityServiceMock()
		localityHandler := handler.NewLocalityHandler(localityService)
		localityService.On("Create", mock.Anything).Return(model.Locality{}, eh.GetErrInternalServer(eh.LOCALITY))

		body, err := json.Marshal(localityRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/localities", localityHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/localities", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{"status": "Internal Server Error","message": "internal server error"}`
		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
		localityService.AssertExpectations(t)
	})
}
