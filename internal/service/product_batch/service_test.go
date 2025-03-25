package service_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	productRepo "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	productBatchRepo "github.com/luisantonisu/wave15-grupo4/internal/repository/product_batch"
	sectionRepo "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/product_batch"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func TestProductBatchService_Create(t *testing.T) {
	productID := 1
	sectionID := 10

	input := model.ProductBatchAttributes{
		ProductID: productID,
		SectionID: sectionID,
	}

	expected := model.ProductBatch{
		ID:                     1,
		ProductBatchAttributes: input,
	}

	t.Run("caso 1: creación exitosa", func(t *testing.T) {
		mockProductRepo := productRepo.NewMockRepository()
		mockSectionRepo := sectionRepo.NewMockRepository()
		mockBatchRepo := productBatchRepo.NewMockProductBatchRepository()

		service := service.NewProductBatchService(mockBatchRepo, mockSectionRepo, mockProductRepo)

		mockProductRepo.On("GetProductByID", productID).Return(model.Product{ID: productID}, nil)
		mockSectionRepo.On("GetByID", sectionID).Return(model.Section{ID: sectionID}, nil)
		mockBatchRepo.On("Create", input).Return(expected, nil)

		result, err := service.Create(input)

		require.NoError(t, err)
		require.Equal(t, expected, result)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
		mockBatchRepo.AssertExpectations(t)
	})

	t.Run("caso 2: producto no existe", func(t *testing.T) {
		mockProductRepo := productRepo.NewMockRepository()
		mockSectionRepo := sectionRepo.NewMockRepository()
		mockBatchRepo := productBatchRepo.NewMockProductBatchRepository()

		service := service.NewProductBatchService(mockBatchRepo, mockSectionRepo, mockProductRepo)

		mockProductRepo.On("GetProductByID", productID).Return(model.Product{}, eh.GetErrNotFound(eh.PRODUCT))

		result, err := service.Create(input)

		require.Error(t, err)
		require.Equal(t, eh.GetErrForeignKey(eh.PRODUCT), err)
		require.Equal(t, model.ProductBatch{}, result)

		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertNotCalled(t, "GetByID", mock.Anything)
		mockBatchRepo.AssertNotCalled(t, "Create")
	})

	t.Run("caso 3: sección no existe", func(t *testing.T) {
		mockProductRepo := productRepo.NewMockRepository()
		mockSectionRepo := sectionRepo.NewMockRepository()
		mockBatchRepo := productBatchRepo.NewMockProductBatchRepository()

		service := service.NewProductBatchService(mockBatchRepo, mockSectionRepo, mockProductRepo)

		mockProductRepo.On("GetProductByID", productID).Return(model.Product{ID: productID}, nil)
		mockSectionRepo.On("GetByID", sectionID).Return(model.Section{}, eh.GetErrNotFound(eh.SECTION))

		result, err := service.Create(input)

		require.Error(t, err)
		require.Equal(t, eh.GetErrForeignKey(eh.SECTION), err)
		require.Equal(t, model.ProductBatch{}, result)

		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
		mockBatchRepo.AssertNotCalled(t, "Create")
	})
}
