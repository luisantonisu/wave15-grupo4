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
	service "github.com/luisantonisu/wave15-grupo4/internal/service/section"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
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

func TestSectionHandler_Create(t *testing.T) {

	t.Run("case 1: create section successfully", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Create", mock.Anything).Return(mockSection, nil)

		body, err := json.Marshal(sectionRequest)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sections", sectionHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/sections", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": {"id":1,"section_number":"1","current_temperature":5.5,"minimum_temperature":5.5,"current_capacity":5,"minimum_capacity":5,"maximum_capacity":5,"warehouse_id":5,"product_type_id":5}
		}`

		require.Equal(t, http.StatusCreated, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 2: invalid data - section missing fields", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Create", mock.Anything).Return(model.Section{}, eh.GetErrInvalidData(eh.SECTION))

		body, err := json.Marshal(dto.SectionRequestDTO{
			SectionNumber:      nil,
			CurrentTemperature: nil,
			MinimumTemperature: nil,
			CurrentCapacity:    nil,
			MinimumCapacity:    &minimumCapacity,
			MaximumCapacity:    &maximumCapacity,
			WarehouseID:        &warehouseID,
			ProductTypeID:      &productTypeID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sections", sectionHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/sections", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Unprocessable Entity",
			"message": "invalid data: section"
		}`
		require.Equal(t, http.StatusUnprocessableEntity, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 3: conflict - section number duplicated", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Create", mock.Anything).Return(model.Section{}, eh.GetErrAlreadyExistsCompose(eh.SECTION, eh.SECTION_NUMBER))

		body, err := json.Marshal(dto.SectionRequestDTO{
			SectionNumber:      &sectionNumber,
			CurrentTemperature: &currentTemperature,
			MinimumTemperature: &minimumTemperature,
			CurrentCapacity:    &currentCapacity,
			MinimumCapacity:    &minimumCapacity,
			MaximumCapacity:    &maximumCapacity,
			WarehouseID:        &warehouseID,
			ProductTypeID:      &productTypeID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sections", sectionHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/sections", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Conflict",
			"message": "section with that section number already exists"
		}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 4: not found - warehouse id doesn't exist", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Create", mock.Anything).Return(model.Section{}, eh.GetErrForeignKey(eh.WAREHOUSE))

		body, err := json.Marshal(dto.SectionRequestDTO{
			SectionNumber:      &sectionNumber,
			CurrentTemperature: &currentTemperature,
			MinimumTemperature: &minimumTemperature,
			CurrentCapacity:    &currentCapacity,
			MinimumCapacity:    &minimumCapacity,
			MaximumCapacity:    &maximumCapacity,
			WarehouseID:        &warehouseID,
			ProductTypeID:      &productTypeID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sections", sectionHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/sections", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Conflict",
			"message": "warehouse foreign key not found"
		}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 5: not found - product id doesn't exist", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Create", mock.Anything).Return(model.Section{}, eh.GetErrForeignKey(eh.PRODUCT))

		body, err := json.Marshal(dto.SectionRequestDTO{
			SectionNumber:      &sectionNumber,
			CurrentTemperature: &currentTemperature,
			MinimumTemperature: &minimumTemperature,
			CurrentCapacity:    &currentCapacity,
			MinimumCapacity:    &minimumCapacity,
			MaximumCapacity:    &maximumCapacity,
			WarehouseID:        &warehouseID,
			ProductTypeID:      &productTypeID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sections", sectionHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/sections", bytes.NewReader(body)), httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Conflict",
			"message": "product foreign key not found"
		}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 6: bad request - invalid body", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Create", mock.Anything).Return(model.Section{}, eh.INVALID_BODY)

		body, err := json.Marshal(`{
			"SectionNumber": "1",
			"CurrentTemperature": "5.5",
			"MinimumTemperature": "5.5",
			"CurrentCapacity": 5,
		}`)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Post("/sections", sectionHandler.Create())

		// Act
		req, res := httptest.NewRequest(http.MethodPost, "/sections", bytes.NewReader(body)), httptest.NewRecorder()
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

	t.Run("case 3: get section by id successfully", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("GetByID", 1).Return(mockSection, nil)

		rt := chi.NewRouter()
		rt.Get("/sections/{id}", sectionHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": {"id" : 1,
					  "section_number" : "1",
					  "current_temperature" : 5.5,
					  "minimum_temperature" : 5.5,
					  "current_capacity" : 5,
					  "minimum_capacity" : 5,
					  "maximum_capacity" : 5,
					  "warehouse_id" : 5,
					  "product_type_id" : 5}
			}`

		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 4: not found - section id doesn't exist", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("GetByID", 2).Return(model.Section{}, eh.GetErrNotFound(eh.SECTION))

		rt := chi.NewRouter()
		rt.Get("/sections/{id}", sectionHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Not Found",
			"message": "section not found"
		}`

		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 5: bad request - invalid section id", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)

		rt := chi.NewRouter()
		rt.Get("/sections/{id}", sectionHandler.GetByID())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/abc", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`

		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

}

func TestSectionHandler_Update(t *testing.T) {

	t.Run("case 1: update section successfully", func(t *testing.T) {
		//Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Patch", 1, mock.Anything).Return(mockSection, nil)

		body, err := json.Marshal(model.SectionAttributes{
			SectionNumber: &sectionNumber,
			WarehouseID:   &warehouseID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sections/{id}", sectionHandler.Patch())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sections/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": {"id":1,"section_number":"1","current_temperature":5.5,"minimum_temperature":5.5,"current_capacity":5,"minimum_capacity":5,"maximum_capacity":5,"warehouse_id":5,"product_type_id":5}
		}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 2: not found - section id doesn't exist", func(t *testing.T) {
		//Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Patch", 2, mock.Anything).Return(model.Section{}, eh.GetErrNotFound(eh.SECTION))

		body, err := json.Marshal(model.SectionAttributes{
			SectionNumber: &sectionNumber,
			WarehouseID:   &warehouseID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sections/{id}", sectionHandler.Patch())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sections/2", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Not Found",
			"message": "section not found"
		}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 3: bad request - section invalid id", func(t *testing.T) {
		//Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Patch", "AA", mock.Anything).Return(model.Section{}, eh.INVALID_ID)

		body, err := json.Marshal(model.SectionAttributes{
			SectionNumber: &sectionNumber,
			WarehouseID:   &warehouseID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sections/{id}", sectionHandler.Patch())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sections/AA", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 4: invalid data - invalid request body", func(t *testing.T) {
		//Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Patch", 1, mock.Anything).Return(model.Section{}, eh.INVALID_BODY)

		body, err := json.Marshal(`{
			"SectionNumber": "1",
			"WarehouseID": 5
		}`)
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sections/{id}", sectionHandler.Patch())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sections/1", bytes.NewReader(body)), httptest.NewRecorder()
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

	t.Run("case 5: foreign key - warehouse id doesn't exist", func(t *testing.T) {
		//Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Patch", 1, mock.Anything).Return(model.Section{}, eh.GetErrForeignKey(eh.WAREHOUSE))

		body, err := json.Marshal(model.SectionAttributes{
			SectionNumber: &sectionNumber,
			WarehouseID:   &warehouseID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sections/{id}", sectionHandler.Patch())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sections/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Conflict",
			"message": "warehouse foreign key not found"
		}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 6: foreign key - product id doesn't exist", func(t *testing.T) {
		//Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Patch", 1, mock.Anything).Return(model.Section{}, eh.GetErrForeignKey(eh.PRODUCT))

		body, err := json.Marshal(model.SectionAttributes{
			SectionNumber: &sectionNumber,
			WarehouseID:   &warehouseID,
		})
		require.NoError(t, err)

		rt := chi.NewRouter()
		rt.Patch("/sections/{id}", sectionHandler.Patch())

		// Act
		req, res := httptest.NewRequest(http.MethodPatch, "/sections/1", bytes.NewReader(body)), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Conflict",
			"message": "product foreign key not found"
		}`
		require.Equal(t, http.StatusConflict, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestSectionHandler_Delete(t *testing.T) {
	t.Run("case 1: delete section successfully", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Delete", 1).Return(nil)

		rt := chi.NewRouter()
		rt.Delete("/sections/{id}", sectionHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/sections/1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		require.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("case 2: not found - section id doesn't exist", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Delete", 2).Return(eh.GetErrNotFound(eh.SECTION))

		rt := chi.NewRouter()
		rt.Delete("/sections/{id}", sectionHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/sections/2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expected := `{
			"status": "Not Found",
			"message": "section not found"
		}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.JSONEq(t, expected, res.Body.String())
	})

	t.Run("case 3: bad request - section invalid id", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)
		sectionService.On("Delete", "AA").Return(eh.INVALID_ID)

		rt := chi.NewRouter()
		rt.Delete("/sections/{id}", sectionHandler.Delete())

		// Act
		req, res := httptest.NewRequest(http.MethodDelete, "/sections/AA", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expected := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.JSONEq(t, expected, res.Body.String())
	})
}

func TestSectionHandler_Report(t *testing.T) {
	t.Run("case 1: get report by section id successfully", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)

		sectionID := 1
		sectionService.On("Report", &sectionID).Return([]model.ReportProductsBatches{
			{
				SectionID:     1,
				SectionNumber: 101,
				ProductsCount: 5,
			},
		}, nil)

		rt := chi.NewRouter()
		rt.Get("/sections/reportProductsBatches", sectionHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/reportProductsBatches?id=1", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": [{"section_id":1, "section_number":101, "products_count":5}]
		}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 2: get all sections report successfully", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)

		// Cambiamos -1 por nil, porque queremos que traiga todas las secciones
		sectionService.On("Report", (*int)(nil)).Return([]model.ReportProductsBatches{
			{
				SectionID:     1,
				SectionNumber: 101,
				ProductsCount: 5,
			},
			{
				SectionID:     2,
				SectionNumber: 202,
				ProductsCount: 10,
			},
		}, nil)

		rt := chi.NewRouter()
		rt.Get("/sections/reportProductsBatches", sectionHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/reportProductsBatches", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"data": [
				{"section_id":1, "section_number":101, "products_count":5},
				{"section_id":2, "section_number":202, "products_count":10}
			]
		}`
		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 3: bad request - invalid section id", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)

		rt := chi.NewRouter()
		rt.Get("/sections/reportProductsBatches", sectionHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/reportProductsBatches?id=hi", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		require.Equal(t, http.StatusBadRequest, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("case 4: not found - section id doesn't exist", func(t *testing.T) {
		// Arrange
		sectionService := service.NewSectionMock()
		sectionHandler := NewSectionHandler(sectionService)

		badSectionID := 2
		sectionService.On("Report", &badSectionID).Return([]model.ReportProductsBatches{}, eh.GetErrNotFound(eh.SECTION))

		rt := chi.NewRouter()
		rt.Get("/sections/reportProductsBatches", sectionHandler.Report())

		// Act
		req, res := httptest.NewRequest(http.MethodGet, "/sections/reportProductsBatches?id=2", nil), httptest.NewRecorder()
		rt.ServeHTTP(res, req)

		// Assert
		expectedBody := `{
			"status": "Not Found",
			"message": "section not found"
		}`
		require.Equal(t, http.StatusNotFound, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})
}
