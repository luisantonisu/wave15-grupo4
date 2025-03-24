package service

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	productRepo "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	productServ "github.com/luisantonisu/wave15-grupo4/internal/repository/product_record"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	LastUpdateDate = "2021-09-01"
	PurchasePrice  = 10.0
	SalePrice      = 20.0
	ProductId      = 1
	BadProductId   = 2
	mockProduct    = model.Product{ID: 1, ProductAttributes: model.ProductAttributes{ProductCode: new(string)}} // Product mock
)

func TestProductRecordService_Create(t *testing.T) {
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
		mockProductRepo := productRepo.NewMockRepository()
		mockProductRecordRepo := productServ.NewMockReportRepository()
		service := NewProductRecordService(mockProductRecordRepo, mockProductRepo)
		// Given

		mockProductRepo.On("GetProductByID", 1).Return(mockProduct, nil)

		mockProductRecordRepo.On("CreateProductRecord", mock.Anything).Return(expectedProductRecord, nil)
		pr, err := service.CreateProductRecord(productRecordAttributes)

		require.NoError(t, err)
		require.Equal(t, expectedProductRecord, pr)
		mockProductRecordRepo.AssertExpectations(t)
	})
	t.Run("case 2: bad product record request/creation", func(t *testing.T) {
		mockProductRepo := productRepo.NewMockRepository()
		mockProductRecordRepo := productServ.NewMockReportRepository()
		service := NewProductRecordService(mockProductRecordRepo, mockProductRepo)
		// Given

		// Mock GetProductByID to return an error
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT))

		// No need to mock CreateProductRecord since it won't be called
		productRecordAttributes := model.ProductRecordAtrributes{
			LastUpdateDate: &LastUpdateDate,
			PurchasePrice:  &PurchasePrice,
			SalePrice:      &SalePrice,
			ProductId:      &ProductId,
		}

		// Act
		pr, err := service.CreateProductRecord(productRecordAttributes)

		// Assert
		require.Error(t, err)
		require.Equal(t, model.ProductRecord{}, pr)
		mockProductRepo.AssertExpectations(t)
		mockProductRecordRepo.AssertExpectations(t)
	})

}

func TestMockProductRecordService_CreateProductRecord(t *testing.T) {
	// Arrange
	mockService := NewMockReportService()

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

	mockService.On("CreateProductRecord", mock.Anything).Return(expectedProductRecord, nil)

	// Act
	result, err := mockService.CreateProductRecord(productRecordAttributes)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedProductRecord, result)
	mockService.AssertExpectations(t)
}
