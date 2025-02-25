package service

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/section"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	SectionNumber      = "1"
	CurrentTemperature = 5
	MinimumTemperature = 5
	CurrentCapacity    = 5
	MinimumCapacity    = 5
	MaximumCapacity    = 5
	WarehouseID        = 5
	ProductTypeID      = 5
	MockSection        = model.Section{ID: 1,
		SectionAttributes: model.SectionAttributes{
			SectionNumber: &SectionNumber,
		}}
)

func TestSectionService_Create(t *testing.T) {

}

func TestSectionService_Read(t *testing.T) {
	mockRepo := repository.NewMockRepository()
	service := NewSectionService(mockRepo)

	t.Run("Case 1: find all", func(t *testing.T) {
		mockSections := []model.Section{
			{
				ID: 1,
				SectionAttributes: model.SectionAttributes{
					SectionNumber: &SectionNumber,
				},
			},
		}
		mockRepo.On("GetAll").Return(mockSections, nil)

		sections, err := service.GetAll()

		require.NoError(t, err)
		require.Equal(t, mockSections, sections)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Case 2: find by id non existent", func(t *testing.T) {
		mockRepo.On("GetByID", 2).Return(model.Section{}, errorHandler.GetErrNotFound(errorHandler.SECTION))

		section, err := service.GetByID(2)
		require.Error(t, err)
		require.Equal(t, model.Section{}, section)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Case 3: find by id existent", func(t *testing.T) {
		mockRepo.On("GetByID", 1).Return(MockSection, nil)

		section, err := service.GetByID(1)
		require.NoError(t, err)
		require.Equal(t, MockSection, section)
		mockRepo.AssertExpectations(t)
	})

}
