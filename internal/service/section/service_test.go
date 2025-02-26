package service

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	CurrentTemperature = 5.5
	MinimumTemperature = 5.5
	CurrentCapacity    = 5
	MinimumCapacity    = 5
	MaximumCapacity    = 5
	WarehouseID        = 5
	ProductTypeID      = 5

	MockSectionAtributes = model.SectionAttributes{CurrentTemperature: &CurrentTemperature, MinimumTemperature: &MinimumTemperature, CurrentCapacity: &CurrentCapacity, MinimumCapacity: &MinimumCapacity, MaximumCapacity: &MaximumCapacity, WarehouseID: &WarehouseID, ProductTypeID: &ProductTypeID}
	MockSection          = model.Section{ID: 1, SectionAttributes: MockSectionAtributes}
	MockSections         = []model.Section{
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
	mockRepo := repository.NewMockRepository()
	service := NewSectionService(mockRepo)

	t.Run("Case 1: create section successfully", func(t *testing.T) {
		t.Cleanup(func() { mockRepo.ExpectedCalls = nil })
		mockRepo.On("Create", MockSectionAtributes).Return(MockSection, nil)

		section, err := service.Create(MockSectionAtributes)

		require.NoError(t, err)
		require.NotNil(t, section)
		require.Equal(t, MockSection, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Case 2: conflict - section number id already exist", func(t *testing.T) {
		errId := eh.GetErrAlreadyExists(eh.SECTION_NUMBER)
		mockRepo.On("Create", MockSectionAtributes).Return(model.Section{}, errId)

		section, err := service.Create(MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockRepo.AssertExpectations(t)
	})
}

func TestSectionService_Read(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewSectionService(mockRepo)

	t.Run("Case 1: get all sections successfully", func(t *testing.T) {
		t.Cleanup(func() { mockRepo.ExpectedCalls = nil })
		mockRepo.On("GetAll").Return(MockSections, nil)

		sections, err := service.GetAll()

		require.NoError(t, err)
		require.NotNil(t, sections)
		require.Equal(t, MockSections, sections)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Case 2: get all sections successfully, no sections found", func(t *testing.T) {
		t.Cleanup(func() { mockRepo.ExpectedCalls = nil })
		mockRepo.On("GetAll").Return([]model.Section{}, nil)

		section, err := service.GetAll()

		require.NoError(t, err)
		require.Empty(t, section)
		require.Equal(t, []model.Section{}, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Case 3: find by id non existent", func(t *testing.T) {
		errId := eh.GetErrNotFound(eh.SECTION)
		mockRepo.On("GetByID", 3).Return(model.Section{}, errId)

		section, err := service.GetByID(3)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Case 4: find by id existent", func(t *testing.T) {
		mockRepo.On("GetByID", 1).Return(MockSection, nil)

		section, err := service.GetByID(1)

		require.NoError(t, err)
		require.NotNil(t, section)
		require.Equal(t, MockSection, section)
		mockRepo.AssertExpectations(t)
	})

}

func TestSectionService_Update(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewSectionService(mockRepo)

	t.Run("Case 1: update section successfully", func(t *testing.T) {
		t.Cleanup(func() { mockRepo.ExpectedCalls = nil })
		mockRepo.On("Patch", 1, MockSectionAtributes).Return(MockSection, nil)

		section, err := service.Patch(1, MockSectionAtributes)

		require.NoError(t, err)
		require.NotNil(t, section)
		require.Equal(t, MockSection, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Case 2: not found - update section, id non existent", func(t *testing.T) {
		errId := eh.GetErrNotFound(eh.SECTION)
		mockRepo.On("Patch", 1, MockSectionAtributes).Return(model.Section{}, errId)

		section, err := service.Patch(1, MockSectionAtributes)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		require.Equal(t, model.Section{}, section)
		mockRepo.AssertExpectations(t)
	})
}

func TestSectionService_Delete(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewSectionService(mockRepo)

	t.Run("Case 1:", func(t *testing.T) {
		t.Cleanup(func() { mockRepo.ExpectedCalls = nil })
		mockRepo.On("Delete", 1).Return(nil)

		err := service.Delete(1)

		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Case 2:", func(t *testing.T) {
		errId := eh.GetErrNotFound(eh.BUYER)
		mockRepo.On("Delete", 1).Return(errId)

		err := service.Delete(1)

		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		mockRepo.AssertExpectations(t)
	})
}
