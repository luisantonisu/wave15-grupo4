package service

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	productID                      = 1
	productCode                    = "P001"
	description                    = "Test Product"
	width                          = 10.0
	height                         = 10.0
	length                         = 10.0
	netWeight                      = 10.0
	expirationRate                 = 10.0
	recommendedFreezingTemperature = -18.0
	freezingRate                   = 10.0
	productTypeId                  = 1
	sellerId                       = 1
	mockProduct                    = model.Product{ID: 1, ProductAttributes: model.ProductAttributes{ProductCode: new(string)}}
	mockRecord                     = model.ProductRecordCount{ProductID: 1, Description: "Record", Count: 1}
)

func TestValueCheck(t *testing.T) {
	productAttributes := model.ProductAttributes{
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
	}
	err := ValueCheck(productAttributes)
	require.NoError(t, err)
	t.Run("Invalid product code", func(t *testing.T) {
		productAttributes.ProductCode = new(string)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid description", func(t *testing.T) {
		productAttributes.ProductCode = &productCode
		productAttributes.Description = new(string)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid width", func(t *testing.T) {
		productAttributes.Description = &description
		productAttributes.Width = new(float64)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid height", func(t *testing.T) {
		productAttributes.Width = &width
		productAttributes.Height = new(float64)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid length", func(t *testing.T) {
		productAttributes.Height = &height
		productAttributes.Length = new(float64)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid net weight", func(t *testing.T) {
		productAttributes.Length = &length
		productAttributes.NetWeight = new(float64)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid expiration rate", func(t *testing.T) {
		productAttributes.NetWeight = &netWeight
		productAttributes.ExpirationRate = new(float64)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid recommended freezing temperature", func(t *testing.T) {
		productAttributes.ExpirationRate = &expirationRate
		productAttributes.RecommendedFreezingTemperature = nil
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid freezing rate", func(t *testing.T) {
		productAttributes.RecommendedFreezingTemperature = &recommendedFreezingTemperature
		productAttributes.FreezingRate = new(float64)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid product type ID", func(t *testing.T) {
		productAttributes.FreezingRate = &freezingRate
		productAttributes.ProductTypeID = new(int)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
	t.Run("Invalid seller ID", func(t *testing.T) {
		productAttributes.ProductTypeID = &productTypeId
		productAttributes.SellerID = new(int)
		err := ValueCheck(productAttributes)
		require.Error(t, err)
	})
}

func TestGetProduct(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewProductService(mockRepo)

	mockProducts := map[int]model.Product{
		1: {ID: 1, ProductAttributes: model.ProductAttributes{ProductCode: new(string)}},
	}
	mockRepo.On("GetProduct").Return(mockProducts, nil)

	products, err := service.GetProduct()
	require.NoError(t, err)
	require.Equal(t, mockProducts, products)
	mockRepo.AssertExpectations(t)
}

func TestGetProductByID(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewProductService(mockRepo)

	t.Run("Get product by ID success", func(t *testing.T) {
		mockRepo.On("GetProductByID", 1).Return(mockProduct, nil)

		product, err := service.GetProductByID(1)
		require.NoError(t, err)
		require.Equal(t, mockProduct, product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Get product by ID not found", func(t *testing.T) {
		mockRepo.On("GetProductByID", 2).Return(model.Product{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT))

		product, err := service.GetProductByID(2)

		require.Error(t, err)
		require.Equal(t, model.Product{}, product)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProductRecord(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewProductService(mockRepo)

	mockRecords := map[int]model.ProductRecordCount{
		1: {ProductID: 1, Description: "Record 1", Count: 10},
	}
	mockRepo.On("GetProductRecord").Return(mockRecords, nil)

	records, err := service.GetProductRecord()
	require.NoError(t, err)
	require.Equal(t, mockRecords, records)
	mockRepo.AssertExpectations(t)
}

func TestGetProductRecordByID(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewProductService(mockRepo)

	mockRepo.On("GetProductRecordByID", 1).Return(mockRecord, nil)

	record, err := service.GetProductRecordByID(1)
	require.NoError(t, err)
	require.Equal(t, mockRecord, record)
	mockRepo.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {

	// mockProduct := model.Product{ID: 1, ProductAttributes: model.ProductAttributes{ProductCode: new(string)}}
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
	productAttributes := &model.ProductAttributes{
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
	}

	t.Run("Create product success", func(t *testing.T) {
		mockRepo := repository.NewMockRepository()
		service := NewProductService(mockRepo)
		mockRepo.On("CreateProduct", productAttributes).Return(mockProduct, nil)

		product, err := service.CreateProduct(productAttributes)
		require.NoError(t, err)
		require.Equal(t, mockProduct, product)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Create product error", func(t *testing.T) {
		mockRepo := repository.NewMockRepository()
		service := NewProductService(mockRepo)
		mockRepo.On("CreateProduct", productAttributes).Return(model.Product{}, errorHandler.GetErrInvalidData(errorHandler.PRODUCT))

		product, err := service.CreateProduct(productAttributes)
		require.Error(t, err)
		require.Equal(t, model.Product{}, product)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewProductService(mockRepo)
	t.Run("Delete product success", func(t *testing.T) {
		mockRepo.On("DeleteProduct", 1).Return(nil)

		err := service.DeleteProduct(1)
		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Delete product error", func(t *testing.T) {
		mockRepo.On("DeleteProduct", 2).Return(errorHandler.GetErrNotFound(errorHandler.PRODUCT))

		err := service.DeleteProduct(2)
		require.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

}

func TestUpdateProduct(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewProductService(mockRepo)

	mockProduct.ProductAttributes.Description = new(string)
	productAttributes := &model.ProductAttributes{ProductCode: new(string)}
	t.Run("Update product success", func(t *testing.T) {
		mockRepo.On("UpdateProduct", 1, productAttributes).Return(&mockProduct, nil)

		product, err := service.UpdateProduct(1, productAttributes)
		require.NoError(t, err)
		require.Equal(t, &mockProduct, product)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Update product error", func(t *testing.T) {
		mockRepo.On("UpdateProduct", 4, productAttributes).Return(&model.Product{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT))

		product, err := service.UpdateProduct(4, productAttributes)
		require.Error(t, err)
		require.Equal(t, &model.Product{}, product)
		mockRepo.AssertExpectations(t)
	})
}
