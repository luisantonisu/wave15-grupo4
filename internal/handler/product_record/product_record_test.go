package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	productProductRecordServ "github.com/luisantonisu/wave15-grupo4/internal/service/product_record"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	LastUpdateDate = "2021-09-01"
	PurchasePrice  = 10.0
	SalePrice      = 20.0
	ProductId      = 1
	BadProductId   = 2
)

func TestProductHandler_Create(t *testing.T) {
	productRecordAttributes := model.ProductRecordAtrributes{
		LastUpdateDate: &LastUpdateDate,
		PurchasePrice:  &PurchasePrice,
		SalePrice:      &SalePrice,
		ProductId:      &ProductId,
	}

	expectedProductRecord := model.ProductRecord{
		ID:                      1,
		ProductRecordAtrributes: productRecordAttributes,
	}

	t.Run("case 1: create product record successfully", func(t *testing.T) {
		mockService := productProductRecordServ.NewMockReportService()
		testHandler := NewProductRecordHandler(mockService)

		mockService.On("CreateProductRecord", mock.Anything).Return(expectedProductRecord, nil)

		body, err := json.Marshal(productRecordAttributes)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/productRecords", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Post("/productRecords", testHandler.Create())
		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusCreated, rr.Code)
		expected := `{
					"data":	{
						"id": 1,
						"last_update_date":"2021-09-01",
						"product_id":1, "purchase_price":10,
						"sale_price":20
						}, 
					"message": "Product record created"
					}`
		require.JSONEq(t, expected, rr.Body.String())
	})
	t.Run("case 2: bad product record request/body", func(t *testing.T) {
		mockService := productProductRecordServ.NewMockReportService()
		testHandler := NewProductRecordHandler(mockService)

		body, err := json.Marshal(`{
		LastUpdateDate: 1,
		PurchasePrice:  &PurchasePrice,
		SalePrice:      &SalePrice,
		ProductId:      &ProductId,
		}`)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/productRecords", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Post("/productRecords", testHandler.Create())
		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
		expected := `{"status": "Bad Request","message":"invalid request body"}`
		require.JSONEq(t, expected, rr.Body.String())
	})

	t.Run("case 3: bad product record service", func(t *testing.T) {
		mockService := productProductRecordServ.NewMockReportService()
		testHandler := NewProductRecordHandler(mockService)

		mockService.On("CreateProductRecord", mock.Anything).Return(model.ProductRecord{}, errorHandler.GetErrForeignKey(errorHandler.PRODUCT))

		body, err := json.Marshal(productRecordAttributes)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/productRecords", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Post("/productRecords", testHandler.Create())
		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusConflict, rr.Code)
		expected := `"product foreign key not found"`
		require.JSONEq(t, expected, rr.Body.String())
	})
}
