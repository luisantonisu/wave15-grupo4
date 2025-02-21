package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/product"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productID                      = 1
	productCode                    = "P001"
	description                    = "Test Product"
	updatedDescription             = "Updated Product"
	width                          = 10.0
	height                         = 10.0
	length                         = 10.0
	netWeight                      = 10.0
	expirationRate                 = 10.0
	recommendedFreezingTemperature = -18.0
	freezingRate                   = 10.0
	productTypeId                  = 1
	sellerId                       = 1
)

func TestGetAll(t *testing.T) {
	mockService := new(service.MockProductService)
	testHandler := NewProductHandler(mockService)

	// Mock the service response
	mockService.On("GetProduct").Return(map[int]model.Product{}, nil)

	req, err := http.NewRequest("GET", "/products", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	testHandler.GetAll().ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	expected := `{"message":"success","data":[]}`
	assert.JSONEq(t, expected, rr.Body.String())
}

func TestGetByID(t *testing.T) {
	mockService := new(service.MockProductService)
	testHandler := NewProductHandler(mockService)

	// Mock the service response
	productID := 1
	mockProduct := model.Product{
		ID: productID,
		ProductAttributes: model.ProductAttributes{
			ProductCode:                    &productCode,
			Description:                    &description,
			Width:                          &width,
			Height:                         &height,
			Length:                         &length,
			NetWeight:                      &netWeight,
			ExpirationRate:                 &expirationRate,
			RecommendedFreezingTemperature: &recommendedFreezingTemperature,
			FreezingRate:                   &freezingRate,
			ProductTypeID:                  &productTypeId,
			SellerID:                       &sellerId,
		}}
	// Test case: Get a specific record when `id` is provided
	t.Run("GetProductByID", func(t *testing.T) {
		// Mock the service response
		mockService.On("GetProductByID", productID).Return(mockProduct, nil)

		req, err := http.NewRequest("GET", "/products/"+strconv.Itoa(productID), nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get("/products/{id}", testHandler.GetByID())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		expected := `{
        "message": "success",
        "data": {
            "id": 1,
            "product_code": "P001",
            "description": "Test Product",
            "width": 10.0,
            "height": 10.0,
            "length": 10.0,
            "net_weight": 10.0,
            "expiration_rate": 10.0,
            "recommended_freezing_temperature": -18.0,
            "freezing_rate": 10.0,
            "product_type_id": 1,
            "seller_id": 1
        }
    }`
		assert.JSONEq(t, expected, rr.Body.String())

	})

	// Test case: Get a specific record when `id` is invalid

	t.Run("InvalidId", func(t *testing.T) {
		// Mock the service response
		mockService.On("GetProductByID", string("d")).Return(mockProduct, nil)

		req, err := http.NewRequest("GET", "/products/d", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get("/products/{id}", testHandler.GetByID())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		expected := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		assert.JSONEq(t, expected, rr.Body.String())

	})

	t.Run("ServiceError", func(t *testing.T) {
		mockService.On("GetProductByID", 2).Return(model.Product{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT))

		req, err := http.NewRequest("GET", "/products/"+strconv.Itoa(2), nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get("/products/{id}", testHandler.GetByID())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		expected := `{
            "message": "product not found",
            "status": "Not Found"
        }`
		assert.JSONEq(t, expected, rr.Body.String())
	})
}

func TestGetRecord(t *testing.T) {
	mockService := new(service.MockProductService)
	testHandler := NewProductHandler(mockService)

	// Test case: Get all records when no `id` is provided
	t.Run("GetAllRecords", func(t *testing.T) {
		// Mock the service response
		mockRecords := map[int]model.ProductRecordCount{
			1: {
				ProductID:   1,
				Description: "Record 1",
				Count:       10,
			},
			2: {
				ProductID:   2,
				Description: "Record 2",
				Count:       20,
			},
		}
		mockService.On("GetProductRecord").Return(mockRecords, nil)

		req, err := http.NewRequest("GET", "/products/reportRecords", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get("/products/reportRecords", testHandler.GetRecord())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		expected := `{
            "message": "success",
            "data": [
                {
                    "product_id": 1,
                    "description": "Record 1",
                    "records_count": 10
                },
                {
                    "product_id": 2,
                    "description": "Record 2",
                    "records_count": 20
                }
            ]
        }`
		assert.JSONEq(t, expected, rr.Body.String())
	})

	// Test case: Get a specific record when `id` is provided
	t.Run("GetRecordByID", func(t *testing.T) {
		// Mock the service response
		mockProductRecord := model.ProductRecordCount{
			ProductID:   1,
			Description: "Record",
			Count:       1,
		}
		mockService.On("GetProductRecordByID", 1).Return(mockProductRecord, nil)

		req, err := http.NewRequest("GET", "/products/reportRecords?id=1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get("/products/reportRecords", testHandler.GetRecord())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		expected := `{
            "message": "success",
            "data": [
                {
                    "product_id": 1,
                    "description": "Record",
                    "records_count": 1
                }
            ]
        }`
		assert.JSONEq(t, expected, rr.Body.String())
	})

	// Test case: Handle invalid `id` query parameter
	t.Run("InvalidID", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/products/reportRecords?id=invalid", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get("/products/reportRecords", testHandler.GetRecord())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		expected := `{
            "message": "invalid id",
            "status": "Bad Request"
        }`
		assert.JSONEq(t, expected, rr.Body.String())
	})

	// Test case: Handle service errors
	t.Run("ServiceError", func(t *testing.T) {
		mockService.On("GetProductRecordByID", 2).Return(model.ProductRecordCount{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT_RECORD))

		req, err := http.NewRequest("GET", "/products/reportRecords?id=2", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get("/products/reportRecords", testHandler.GetRecord())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		expected := `{
            "message": "product record not found",
            "status": "Not Found"
        }`
		assert.JSONEq(t, expected, rr.Body.String())
	})
}

func TestCreate(t *testing.T) {
	mockService := new(service.MockProductService)
	testHandler := NewProductHandler(mockService)

	t.Run("ProductCreated", func(t *testing.T) {
		mockProduct := model.Product{
			ID: productID,
			ProductAttributes: model.ProductAttributes{
				ProductCode:                    &productCode,
				Description:                    &description,
				Width:                          &width,
				Height:                         &height,
				Length:                         &length,
				NetWeight:                      &netWeight,
				ExpirationRate:                 &expirationRate,
				RecommendedFreezingTemperature: &recommendedFreezingTemperature,
				FreezingRate:                   &freezingRate,
				ProductTypeID:                  &productTypeId,
				SellerID:                       &sellerId,
			}}
		productRequest := dto.ProductRequestDTO{
			ProductCode:                    &productCode,
			Description:                    &description,
			Width:                          &width,
			Height:                         &height,
			Length:                         &length,
			NetWeight:                      &netWeight,
			ExpirationRate:                 &expirationRate,
			RecommendedFreezingTemperature: &recommendedFreezingTemperature,
			FreezingRate:                   &freezingRate,
			ProductTypeId:                  &productTypeId,
			SellerId:                       &sellerId,
		}
		mockService.On("CreateProduct", mock.Anything).Return(mockProduct, nil)

		body, err := json.Marshal(productRequest)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Post("/products", testHandler.Create())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		expected := `{
        "message": "Product created",
        "data": {
            "id": 1,
            "product_code": "P001",
            "description": "Test Product",
            "width": 10.0,
            "height": 10.0,
            "length": 10.0,
            "net_weight": 10.0,
            "expiration_rate": 10.0,
            "recommended_freezing_temperature": -18.0,
            "freezing_rate": 10.0,
            "product_type_id": 1,
            "seller_id": 1
        }
    }`
		assert.JSONEq(t, expected, rr.Body.String())
	})

	// t.Run("ServiceError", func(t *testing.T) {
	// 	mockService.On("CreateProduct", mock.Anything).Return(model.Product{}, errorHandler.GetErrInvalidData(errorHandler.PRODUCT))

	// 	body, err := json.Marshal(dto.ProductRequestDTO{
	// 		ProductCode:                    nil,
	// 		Description:                    nil,
	// 		Width:                          nil,
	// 		Height:                         nil,
	// 		Length:                         nil,
	// 		NetWeight:                      nil,
	// 		ExpirationRate:                 nil,
	// 		RecommendedFreezingTemperature: nil,
	// 		FreezingRate:                   nil,
	// 		ProductTypeId:                  nil,
	// 		SellerId:                       nil,
	// 	})
	// 	assert.NoError(t, err)

	// 	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	// 	assert.NoError(t, err)
	// 	req.Header.Set("Content-Type", "application/json")

	// 	rr := httptest.NewRecorder()
	// 	fmt.Println("PRINT LINEEEEEEE------", rr)

	// 	r := chi.NewRouter()
	// 	r.Post("/products", testHandler.Create())
	// 	r.ServeHTTP(rr, req)
	// 	fmt.Println("OMGGGGGG", rr)

	// 	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	// 	expected := `{
	//         "status": "Unprocessable Entity",
	//         "message": "invalid data: product"
	//     }`
	// 	assert.JSONEq(t, expected, rr.Body.String())
	// })
}

func TestCreateBad(t *testing.T) {
	mockService := new(service.MockProductService)
	testHandler := NewProductHandler(mockService)

	// Mock the service response
	mockService.On("CreateProduct", mock.Anything).Return(model.Product{}, errorHandler.GetErrInvalidData(errorHandler.PRODUCT))

	body, err := json.Marshal(dto.ProductRequestDTO{
		ProductCode:                    nil,
		Description:                    nil,
		Width:                          nil,
		Height:                         nil,
		Length:                         nil,
		NetWeight:                      nil,
		ExpirationRate:                 nil,
		RecommendedFreezingTemperature: nil,
		FreezingRate:                   nil,
		ProductTypeId:                  nil,
		SellerId:                       nil,
	})
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	r := chi.NewRouter()
	r.Post("/products", testHandler.Create())
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	expected := `{
            "status": "Unprocessable Entity",
            "message": "invalid data: product"
        }`
	assert.JSONEq(t, expected, rr.Body.String())

}

func TestDelete(t *testing.T) {
	mockService := new(service.MockProductService)
	testHandler := NewProductHandler(mockService)

	// Mock the service response
	productID := 1
	t.Run("InvalidIdNumber", func(t *testing.T) {
		// Mock the service response

		req, err := http.NewRequest("DELETE", "/products/0", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Delete("/products/{id}", testHandler.Delete())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		expected := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		assert.JSONEq(t, expected, rr.Body.String())

	})

	t.Run("InvalidId", func(t *testing.T) {
		// Mock the service response

		req, err := http.NewRequest("DELETE", "/products/d", nil)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Delete("/products/{id}", testHandler.Delete())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		expected := `{
			"status": "Bad Request",
			"message": "invalid id"
		}`
		assert.JSONEq(t, expected, rr.Body.String())

	})

	t.Run("DeleteProduct", func(t *testing.T) {
		mockService.On("DeleteProduct", productID).Return(nil)

		req, err := http.NewRequest("DELETE", "/products/"+strconv.Itoa(productID), nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Delete("/products/{id}", testHandler.Delete())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)

	})

}

func TestUpdate(t *testing.T) {
	mockService := new(service.MockProductService)
	testHandler := NewProductHandler(mockService)

	// Mock the service response
	mockProduct := &model.Product{
		ID: productID,
		ProductAttributes: model.ProductAttributes{
			ProductCode:                    &productCode,
			Description:                    &updatedDescription,
			Width:                          &width,
			Height:                         &height,
			Length:                         &length,
			NetWeight:                      &netWeight,
			ExpirationRate:                 &expirationRate,
			RecommendedFreezingTemperature: &recommendedFreezingTemperature,
			FreezingRate:                   &freezingRate,
			ProductTypeID:                  &productTypeId,
			SellerID:                       &sellerId,
		}}

	productRequest := dto.ProductRequestDTO{
		ProductCode:                    &productCode,
		Description:                    &updatedDescription,
		Width:                          &width,
		Height:                         &height,
		Length:                         &length,
		NetWeight:                      &netWeight,
		ExpirationRate:                 &expirationRate,
		RecommendedFreezingTemperature: &recommendedFreezingTemperature,
		FreezingRate:                   &freezingRate,
		ProductTypeId:                  &productTypeId,
		SellerId:                       &sellerId,
	}

	t.Run("NotFound", func(t *testing.T) {
		// Mock the service response
		mockService.On("UpdateProduct", 4, mock.Anything).Return(&model.Product{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT))

		body, err := json.Marshal(productRequest)
		assert.NoError(t, err)

		req, err := http.NewRequest("PATCH", "/products/"+strconv.Itoa(4), bytes.NewBuffer(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Patch("/products/{id}", testHandler.Update())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		expected := `{
            "status": "Not Found",
            "message": "product not found"
        }`
		assert.JSONEq(t, expected, rr.Body.String())
	})
	t.Run("UpdateProduct", func(t *testing.T) {
		mockService.On("UpdateProduct", productID, mock.Anything).Return(mockProduct, nil)

		body, err := json.Marshal(productRequest)
		assert.NoError(t, err)

		req, err := http.NewRequest("PATCH", "/products/"+strconv.Itoa(productID), bytes.NewBuffer(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Patch("/products/{id}", testHandler.Update())
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		expected := `{
        "message": "Product updated",
        "data": [
        {
            "id": 1,
            "product_code": "P001",
            "description": "Updated Product",
            "width": 10.0,
            "height": 10.0,
            "length": 10.0,
            "net_weight": 10.0,
            "expiration_rate": 10.0,
            "recommended_freezing_temperature": -18.0,
            "freezing_rate": 10.0,
            "product_type_id": 1,
            "seller_id": 1
           }
    ]
    }`
		assert.JSONEq(t, expected, rr.Body.String())
	})
}

/* {
    "product_code": "122",
    "width": 2.1,
    "height": 0,
    "description": ""
    aaa
} */
