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
	LastUpdateDate             = "2021-09-01"
	PurchasePrice              = 10.0
	SalePrice                  = 20.0
	ProductId                  = 1
	BadProductId               = 2
	mockProduct                = model.Product{ID: 1, ProductAttributes: model.ProductAttributes{ProductCode: new(string)}}
	badProductRecordAttributes = model.ProductRecordAtrributes{
		LastUpdateDate: nil,
		PurchasePrice:  nil,
		SalePrice:      nil,
		ProductId:      nil,
	}
)

/*
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
*/

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
