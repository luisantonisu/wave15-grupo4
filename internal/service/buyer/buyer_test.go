package service_test

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

// Variables for testing
var (
	cardNumberId  = "123456"
	firstName     = "John"
	lastName      = "Doe"
	cardNumberIdB = "654321"
	firstNameB    = "Jane"
	lastNameB     = "Doe"

	buyerAttributes = model.BuyerAttributes{
		CardNumberId: &cardNumberId,
		FirstName:    &firstName,
		LastName:     &lastName,
	}

	buyerAttributesB = model.BuyerAttributes{
		CardNumberId: &cardNumberIdB,
		FirstName:    &firstNameB,
		LastName:     &lastNameB,
	}
	purchaseOrderReport = model.ReportPurchaseOrders{
		ID:                  1,
		CardNumberId:        cardNumberId,
		FirstName:           firstName,
		LastName:            lastName,
		PurchaseOrdersCount: 1,
	}
	purchaseOrderReportB = model.ReportPurchaseOrders{
		ID:                  2,
		CardNumberId:        cardNumberIdB,
		FirstName:           firstNameB,
		LastName:            lastNameB,
		PurchaseOrdersCount: 2,
	}
)

func TestBuyerService_Create(t *testing.T) {
	t.Run("case 1: create buyer successfully", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		buyer := model.Buyer{
			ID:              1,
			BuyerAttributes: buyerAttributes,
		}
		repo.On("Create", buyerAttributes).Return(buyer, nil)

		// Act
		buyerResult, err := buyerService.Create(buyerAttributes)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, buyer)
		require.Equal(t, buyer, buyerResult)
		repo.AssertExpectations(t)
	})

	t.Run("case 2: conflict - card number id already exist", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		errId := eh.GetErrAlreadyExists(eh.CARD_NUMBER)
		repo.On("Create", buyerAttributes).Return(model.Buyer{}, errId)

		// Act
		buyerResult, err := buyerService.Create(buyerAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, errId, err)
		require.Equal(t, model.Buyer{}, buyerResult)
		repo.AssertExpectations(t)
	})
}

func TestBuyerService_GetAll(t *testing.T) {
	t.Run("case 1: get all buyers successfully", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		buyers := []model.Buyer{
			{
				ID:              1,
				BuyerAttributes: buyerAttributes,
			},
			{
				ID:              2,
				BuyerAttributes: buyerAttributesB,
			},
		}
		repo.On("GetAll").Return(buyers, nil)

		// Act
		buyersResult, err := buyerService.GetAll()

		// Assert
		require.NoError(t, err)
		require.NotNil(t, buyers)
		require.Equal(t, buyers, buyersResult)
		repo.AssertExpectations(t)
	})
	t.Run("case 2: no buyers found", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		repo.On("GetAll").Return([]model.Buyer{}, nil)

		// Act
		buyersResult, err := buyerService.GetAll()

		// Assert
		require.NoError(t, err)
		require.Empty(t, buyersResult)
		require.Equal(t, []model.Buyer{}, buyersResult)
		repo.AssertExpectations(t)
	})
}

func TestBuyerService_GetById(t *testing.T) {
	t.Run("case 1: get buyer by ID successfully", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		buyer := model.Buyer{
			ID:              1,
			BuyerAttributes: buyerAttributes,
		}
		repo.On("GetByID", 1).Return(buyer, nil)

		// Act
		buyerResult, err := buyerService.GetByID(1)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, buyer)
		require.Equal(t, buyer, buyerResult)
		repo.AssertExpectations(t)
	})
	t.Run("case 2: not found - get buyer by ID non existent", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		errId := eh.GetErrNotFound(eh.BUYER)
		repo.On("GetByID", 1).Return(model.Buyer{}, errId)

		// Act
		buyerResult, err := buyerService.GetByID(1)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		require.Equal(t, model.Buyer{}, buyerResult)
		repo.AssertExpectations(t)
	})
}

func TestBuyerService_Update(t *testing.T) {
	t.Run("case 1: update buyer successfully", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		buyer := model.Buyer{
			ID:              1,
			BuyerAttributes: buyerAttributes,
		}
		repo.On("Update", 1, buyerAttributes).Return(buyer, nil)

		// Act
		buyerResult, err := buyerService.Update(1, buyerAttributes)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, buyer)
		require.Equal(t, buyer, buyerResult)
		repo.AssertExpectations(t)
	})
	t.Run("case 2: not found - update buyer non existent", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		errId := eh.GetErrNotFound(eh.BUYER)
		repo.On("Update", 1, buyerAttributes).Return(model.Buyer{}, errId)

		// Act
		buyerResult, err := buyerService.Update(1, buyerAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		require.Equal(t, model.Buyer{}, buyerResult)
		repo.AssertExpectations(t)
	})
}

func TestBuyerService_Delete(t *testing.T) {
	t.Run("case 1: delete buyer successfully", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		repo.On("Delete", 1).Return(nil)

		// Act
		err := buyerService.Delete(1)

		// Assert
		require.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("case 2: not found - delete buyer non existent", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		errId := eh.GetErrNotFound(eh.BUYER)
		repo.On("Delete", 1).Return(errId)

		// Act
		err := buyerService.Delete(1)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		repo.AssertExpectations(t)
	})
}

func TestBuyerService_ReportPurchaseOrder(t *testing.T) {
	t.Run("case 1: get purchase order report without id successfully", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		var id *int
		report := []model.ReportPurchaseOrders{
			purchaseOrderReport,
			purchaseOrderReportB,
		}
		repo.On("PurchaseOrderReport", id).Return(report, nil)

		// Act
		reportResult, err := buyerService.PurchaseOrderReport(nil)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, report)
		require.Equal(t, report, reportResult)
		repo.AssertExpectations(t)
	})

	t.Run("case 2: get purchase order report with id successfully", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		id := 1
		report := []model.ReportPurchaseOrders{
			purchaseOrderReport,
		}
		repo.On("PurchaseOrderReport", &id).Return(report, nil)

		// Act
		reportResult, err := buyerService.PurchaseOrderReport(&id)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, report)
		require.Equal(t, report, reportResult)
		repo.AssertExpectations(t)
	})
	t.Run("case 3: get purchase order report with id succesfully but empty", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		id := 1
		repo.On("PurchaseOrderReport", &id).Return([]model.ReportPurchaseOrders{}, nil)

		// Act
		reportResult, err := buyerService.PurchaseOrderReport(&id)

		// Assert
		require.NoError(t, err)
		require.Empty(t, reportResult)
		repo.AssertExpectations(t)
	})
	t.Run("case 4: not found - get purchase order report with id non existent", func(t *testing.T) {
		// Arrange
		repo := repository.NewBuyerRepositoryMock()
		buyerService := service.NewBuyerService(repo)
		id := 1
		errId := eh.GetErrNotFound(eh.BUYER)
		repo.On("PurchaseOrderReport", &id).Return([]model.ReportPurchaseOrders{}, errId)

		// Act
		reportResult, err := buyerService.PurchaseOrderReport(&id)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errId, err)
		require.Empty(t, reportResult)
		repo.AssertExpectations(t)
	})
}
