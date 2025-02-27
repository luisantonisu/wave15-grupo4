package service

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repositoryProduct "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	repositorySection "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	repositoryWarehouse "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	SectionNumber      = "1"
	CurrentTemperature = 5.5
	MinimumTemperature = 5.5
	CurrentCapacity    = 5
	MinimumCapacity    = 5
	MaximumCapacity    = 5
	WarehouseID        = 1
	Product            = 1

	MockSectionAtributesIncompleted = model.SectionAttributes{SectionNumber: nil, CurrentTemperature: &CurrentTemperature, MinimumTemperature: &MinimumTemperature, CurrentCapacity: &CurrentCapacity, MinimumCapacity: &MinimumCapacity, MaximumCapacity: &MaximumCapacity, WarehouseID: &WarehouseID, ProductTypeID: &Product}
	MockSectionAtributes            = model.SectionAttributes{SectionNumber: &SectionNumber, CurrentTemperature: &CurrentTemperature, MinimumTemperature: &MinimumTemperature, CurrentCapacity: &CurrentCapacity, MinimumCapacity: &MinimumCapacity, MaximumCapacity: &MaximumCapacity, WarehouseID: &WarehouseID, ProductTypeID: &Product}
	MockSection                     = model.Section{ID: 1, SectionAttributes: MockSectionAtributes}
	MockSections                    = []model.Section{
		{
			ID:                1,
			SectionAttributes: MockSectionAtributes,
		},
		{
			ID:                2,
			SectionAttributes: MockSectionAtributes,
		},
	}
)

func TestSectionService_Create(t *testing.T) {
	mockSectionRepo := repositorySection.NewMockRepository()
	mockProductRepo := repositoryProduct.NewMockRepository()
	mockWarehouseRepo := repositoryWarehouse.NewWarehouseRepositoryMock()
	service := NewSectionService(mockSectionRepo, mockProductRepo, mockWarehouseRepo)

	t.Run("case 1: create section successfully", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{ID: Product}, nil)
		mockSectionRepo.On("Create", MockSectionAtributes).Return(MockSection, nil)

		section, err := service.Create(MockSectionAtributes)

		require.NoError(t, err)
		require.NotNil(t, section)
		require.Equal(t, MockSection, section)
		mockWarehouseRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 2: invalid data - section missing fields", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		errId := eh.GetErrInvalidData(eh.SECTION)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{ID: Product}, nil)
		mockSectionRepo.On("Create", MockSectionAtributesIncompleted).Return(model.Section{}, errId)

		section, err := service.Create(MockSectionAtributesIncompleted)

		require.Error(t, err)
		require.ErrorIs(t, err, errId)
		require.Equal(t, eh.GetErrInvalidData(eh.SECTION), err)
		require.Equal(t, model.Section{}, section)
		mockWarehouseRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 3: foreign key - warehouse id not found", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		errId := eh.GetErrForeignKey(eh.WAREHOUSE)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{}, errId)

		section, err := service.Create(MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockWarehouseRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 4: foreign key - product id not found", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		errId := eh.GetErrForeignKey(eh.PRODUCT)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{}, errId)

		section, err := service.Create(MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockWarehouseRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 5: conflict - section number id already exist", func(t *testing.T) {
		errId := eh.GetErrAlreadyExists(eh.SECTION_NUMBER)
		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{ID: Product}, nil)
		mockSectionRepo.On("Create", MockSectionAtributes).Return(model.Section{}, errId)

		section, err := service.Create(MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockSectionRepo.AssertExpectations(t)
	})
}

func TestSectionService_Read(t *testing.T) {
	mockSectionRepo := repositorySection.NewMockRepository()
	mockProductRepo := repositoryProduct.NewMockRepository()
	mockWarehouseRepo := repositoryWarehouse.NewWarehouseRepositoryMock()
	service := NewSectionService(mockSectionRepo, mockProductRepo, mockWarehouseRepo)

	t.Run("case 1: get all sections successfully", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		mockSectionRepo.On("GetAll").Return(MockSections, nil)

		sections, err := service.GetAll()

		require.NoError(t, err)
		require.NotNil(t, sections)
		require.Equal(t, MockSections, sections)
		mockSectionRepo.AssertExpectations(t)

	})

	t.Run("case 2: get all sections successfully, no sections found", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		mockSectionRepo.On("GetAll").Return([]model.Section{}, nil)

		section, err := service.GetAll()

		require.NoError(t, err)
		require.Empty(t, section)
		require.Equal(t, []model.Section{}, section)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 3: find by id non existent", func(t *testing.T) {
		errId := eh.GetErrNotFound(eh.SECTION)
		mockSectionRepo.On("GetByID", 3).Return(model.Section{}, errId)

		section, err := service.GetByID(3)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 4: find by id existent", func(t *testing.T) {
		mockSectionRepo.On("GetByID", 1).Return(MockSection, nil)

		section, err := service.GetByID(1)

		require.NoError(t, err)
		require.NotNil(t, section)
		require.Equal(t, MockSection, section)
		mockSectionRepo.AssertExpectations(t)
	})

}

func TestSectionService_Update(t *testing.T) {
	mockSectionRepo := repositorySection.NewMockRepository()
	mockProductRepo := repositoryProduct.NewMockRepository()
	mockWarehouseRepo := repositoryWarehouse.NewWarehouseRepositoryMock()
	service := NewSectionService(mockSectionRepo, mockProductRepo, mockWarehouseRepo)

	t.Run("case 1: update section successfully", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{ID: Product}, nil)
		mockSectionRepo.On("Patch", 1, MockSectionAtributes).Return(MockSection, nil)

		section, err := service.Patch(1, MockSectionAtributes)

		require.NoError(t, err)
		require.NotNil(t, section)
		require.Equal(t, MockSection, section)
		mockWarehouseRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 2: not found - update section, id non existent", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		errId := eh.GetErrNotFound(eh.SECTION)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{ID: Product}, nil)
		mockSectionRepo.On("Patch", 1, MockSectionAtributes).Return(model.Section{}, errId)

		section, err := service.Patch(1, MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 3: invalid data - section missing fields", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		errId := eh.GetErrInvalidData(eh.SECTION)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{ID: Product}, nil)
		mockSectionRepo.On("Patch", 1, MockSectionAtributesIncompleted).Return(model.Section{}, errId)

		section, err := service.Patch(1, MockSectionAtributesIncompleted)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockWarehouseRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 4: foreign key - warehouse id not found", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		errId := eh.GetErrForeignKey(eh.WAREHOUSE)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{}, errId)

		section, err := service.Patch(1, MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockWarehouseRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 5: foreign key - product id not found", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		errId := eh.GetErrForeignKey(eh.PRODUCT)

		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{}, errId)

		section, err := service.Patch(1, MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockWarehouseRepo.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 6: conflict - section number id already exist", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		errId := eh.GetErrAlreadyExists(eh.SECTION_NUMBER)
		mockWarehouseRepo.On("GetByID", 1).Return(model.Warehouse{ID: WarehouseID}, nil)
		mockProductRepo.On("GetProductByID", 1).Return(model.Product{ID: Product}, nil)
		mockSectionRepo.On("Patch", 1, MockSectionAtributes).Return(model.Section{}, errId)

		section, err := service.Patch(1, MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockSectionRepo.AssertExpectations(t)
	})
}

func TestSectionService_Delete(t *testing.T) {
	mockSectionRepo := repositorySection.NewMockRepository()
	mockProductRepo := repositoryProduct.NewMockRepository()
	mockWarehouseRepo := repositoryWarehouse.NewWarehouseRepositoryMock()
	service := NewSectionService(mockSectionRepo, mockProductRepo, mockWarehouseRepo)

	t.Run("case 1: delete section successfully", func(t *testing.T) {
		t.Cleanup(func() {
			mockSectionRepo.ExpectedCalls = nil
			mockProductRepo.ExpectedCalls = nil
			mockWarehouseRepo.ExpectedCalls = nil
		})

		mockSectionRepo.On("Delete", 1).Return(nil)

		err := service.Delete(1)

		require.NoError(t, err)
		mockSectionRepo.AssertExpectations(t)
	})

	t.Run("case 2: not found - section id doesn't exist", func(t *testing.T) {
		errId := eh.GetErrNotFound(eh.BUYER)
		mockSectionRepo.On("Delete", 1).Return(errId)

		err := service.Delete(1)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		mockSectionRepo.AssertExpectations(t)
	})
}
