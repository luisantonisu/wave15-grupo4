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
	productProductRecordServ "github.com/luisantonisu/wave15-grupo4/internal/service/product_record"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
	mockProduct                    = model.Product{
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
	productRequest = dto.ProductRequestDTO{
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
)

func TestProductHandler_Create(t *testing.T) {

	t.Run("case 1: create product successfully", func(t *testing.T) {
		mockService := productProductRecordServ.NewMockReportService()
		testHandler := NewProductRecordHandler(mockService)

		mockService.On("CreateProduct", mock.Anything).Return(mockProduct, nil)

		body, err := json.Marshal(productRequest)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Post("/products", testHandler.Create())
		r.ServeHTTP(rr, req)

		require.Equal(t, http.StatusCreated, rr.Code)
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
		require.JSONEq(t, expected, rr.Body.String())
	})
}
