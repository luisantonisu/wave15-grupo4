package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/section"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	sectionNumber      = "1"
	currentTemperature = 5.5
	minimumTemperature = 5.5
	currentCapacity    = 5
	minimumCapacity    = 5
	maximumCapacity    = 5
	warehouseID        = 5
	productTypeID      = 5

	mockSectionAtributes = model.SectionAttributes{
		SectionNumber:      &sectionNumber,
		CurrentTemperature: &currentTemperature,
		MinimumTemperature: &minimumTemperature,
		CurrentCapacity:    &currentCapacity,
		MinimumCapacity:    &minimumCapacity,
		MaximumCapacity:    &maximumCapacity,
		WarehouseID:        &warehouseID,
		ProductTypeID:      &productTypeID,
	}

	mockSection = model.Section{
		ID:                1,
		SectionAttributes: mockSectionAtributes,
	}

	mockSections = []model.Section{
		{
			ID:                1,
			SectionAttributes: mockSectionAtributes,
		},
	}

	sectionRequest = dto.SectionRequestDTO{
		SectionNumber:      &sectionNumber,
		CurrentTemperature: &currentTemperature,
		MinimumTemperature: &minimumTemperature,
		CurrentCapacity:    &currentCapacity,
		MinimumCapacity:    &minimumCapacity,
		MaximumCapacity:    &maximumCapacity,
		WarehouseID:        &warehouseID,
		ProductTypeID:      &productTypeID,
	}
)

func TestSectionHandler_Get(t *testing.T) {
	t.Run("case 1: get all sections successfully", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("GetAll").Return(mockSections, nil)

		rt := chi.NewRouter()
		rt.Get("/sections", sectionHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": [{"id" : 1,
					  "section_number" : "1",
					  "current_temperature" : 5.5,
					  "minimum_temperature" : 5.5,
					  "current_capacity" : 5,
					  "minimum_capacity" : 5,
					  "maximum_capacity" : 5,
					  "warehouse_id" : 5,
					  "product_type_id" : 5}]
			}`

		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 2: internal server error - get all sections", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("GetAll").Return([]model.Section{}, eh.GetErrDatabase(eh.SECTION))

		rt := chi.NewRouter()
		rt.Get("/sections", sectionHandler.GetAll())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Internal Server Error",
			"message": "database error: section"
		}`

		require.Equal(t, http.StatusInternalServerError, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

}
