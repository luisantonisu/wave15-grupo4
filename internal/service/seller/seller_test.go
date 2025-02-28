package service

import (
	"testing"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	sellerLocality "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"
	sellerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"github.com/stretchr/testify/require"
)

var (
	companyID     = "1234"
	badCompanyID  = "ABC"
	companyName   = "Company X"
	address       = "221B Baker Street"
	telephone     = "123590"
	localityId    = "1"
	badLocalityId = "ABC"

	companyIDB   = "2345"
	companyNameB = "Company Z"
	addressB     = "1007 de Mountain Drive"
	telephoneB   = "2145119"
	localityIdB  = "2"

	sellerAttributes = model.SellerAttributes{
		CompanyID:   &companyID,
		CompanyName: &companyName,
		Address:     &address,
		Telephone:   &telephone,
		LocalityId:  &localityId,
	}

	sellerAttributesB = model.SellerAttributes{
		CompanyID:   new(string),
		CompanyName: new(string),
		Address:     new(string),
		Telephone:   new(string),
		LocalityId:  new(string),
	}

	seller = model.Seller{
		ID: 1,
		SellerAttributes: model.SellerAttributes{
			CompanyID:   &companyID,
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
			LocalityId:  &localityId,
		},
	}
	sellerB = model.Seller{
		ID: 2,
		SellerAttributes: model.SellerAttributes{
			CompanyID:   &companyIDB,
			CompanyName: &companyNameB,
			Address:     &addressB,
			Telephone:   &telephoneB,
			LocalityId:  &localityIdB,
		},
	}
)

func TestValidateSeller(t *testing.T) {

	sellerAttributes := model.SellerAttributes{
		CompanyID:   &companyID,
		CompanyName: &companyName,
		Address:     &address,
		Telephone:   &telephone,
		LocalityId:  &localityId,
	}

	sellerRp := sellerRepository.NewSellerRepositoryMock()
	localityRp := sellerLocality.NewLocalityRepositoryMock()

	sellerSv := NewSellerService(sellerRp, localityRp)

	err := sellerSv.ValidateSeller(sellerAttributes)
	require.NoError(t, err)

	t.Run("case 1: CompanyId is required", func(t *testing.T) {
		sellerAttributes.CompanyID = new(string)
		err := sellerSv.ValidateSeller(sellerAttributes)
		require.Error(t, err)
	})
	t.Run("case 2 CompanyID should contain only numbers", func(t *testing.T) {
		invalidRequest := "Aaa"
		sellerAttributes.CompanyID = &invalidRequest
		err := sellerSv.ValidateSeller(sellerAttributes)
		require.Error(t, err)
	})
	t.Run("case 3 CompanyName is required", func(t *testing.T) {
		sellerAttributes.CompanyID = &companyID
		sellerAttributes.CompanyName = new(string)
		err := sellerSv.ValidateSeller(sellerAttributes)
		require.Error(t, err)
	})
	t.Run("case 4: Address is required", func(t *testing.T) {
		sellerAttributes.CompanyName = &companyName
		sellerAttributes.Address = new(string)
		err := sellerSv.ValidateSeller(sellerAttributes)
		require.Error(t, err)
	})
	t.Run("case 5: Telephone is required", func(t *testing.T) {
		sellerAttributes.Address = &address
		sellerAttributes.Telephone = new(string)
		err := sellerSv.ValidateSeller(sellerAttributes)
		require.Error(t, err)
	})
	t.Run("case 6: LocalityID is required", func(t *testing.T) {
		sellerAttributes.Telephone = &telephone
		sellerAttributes.LocalityId = new(string)
		err := sellerSv.ValidateSeller(sellerAttributes)
		require.Error(t, err)
	})
	t.Run("case 7: LocalityId should contain only numbers", func(t *testing.T) {
		sellerAttributes.Telephone = &telephone
		invalidRequest := "Aaa"
		sellerAttributes.LocalityId = &invalidRequest
		err := sellerSv.ValidateSeller(sellerAttributes)
		require.Error(t, err)
	})
}

func TestSellerService_GetAll(t *testing.T) {
	t.Run("case 1: Find all sellers", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()

		sellerSv := NewSellerService(sellerRp, localityRp)

		var sellersList = []model.Seller{seller, sellerB}

		sellerRp.On("GetAll").Return(sellersList, nil)

		//Act
		sellers, err := sellerSv.GetAll()

		//Assert
		require.NoError(t, err)
		require.Equal(t, sellersList, sellers)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 2: No seller was found", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()

		sellerSv := NewSellerService(sellerRp, localityRp)

		sellerRp.On("GetAll").Return([]model.Seller{}, nil)

		//Act
		sellers, err := sellerSv.GetAll()

		//Assert
		require.NoError(t, err)
		require.Empty(t, sellers)
		require.Equal(t, []model.Seller{}, sellers)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
}

func TestSellerService_GetByID(t *testing.T) {
	t.Run("case 1: get seller by ID successfully", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()

		sellerSv := NewSellerService(sellerRp, localityRp)

		sellerRp.On("GetByID", 1).Return(seller, nil)

		//Act
		sellerResult, err := sellerSv.GetByID(1)

		//Assert
		require.NoError(t, err)
		require.NotNil(t, seller)
		require.Equal(t, seller, sellerResult)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 2:  not found - get seller by ID non existent", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		errSellerNotFound := eh.GetErrNotFound(eh.SELLER)

		sellerRp.On("GetByID", 1).Return(model.Seller{}, errSellerNotFound)

		//Act
		sellerResult, err := sellerSv.GetByID(1)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errSellerNotFound, err)
		require.Equal(t, model.Seller{}, sellerResult)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
}

func TestSellerService_Create(t *testing.T) {
	t.Run("case 1: Create Seller successfully", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)

		localityRp.On("GetByID", 1).Return(model.LocalityDBModel{Id: 1}, nil)
		sellerRp.On("Create", sellerAttributes).Return(seller, nil)

		//Act
		newSeller, err := sellerSv.Create(sellerAttributes)

		//Assert
		require.NoError(t, err)
		require.NotNil(t, seller)
		require.Equal(t, seller, newSeller)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 2: conflict - CID is already exist ", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		errAlreadyExists := eh.GetErrAlreadyExists(eh.SELLER)

		localityRp.On("GetByID", 1).Return(model.LocalityDBModel{Id: 1}, nil)
		sellerRp.On("Create", sellerAttributes).Return(model.Seller{}, errAlreadyExists)

		//Act
		seller, err := sellerSv.Create(sellerAttributes)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrAlreadyExists)
		require.Equal(t, errAlreadyExists, err)
		require.Equal(t, model.Seller{}, seller)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 3: invalid data - seller missing fields", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		errInvalidData := eh.GetErrInvalidData(eh.SELLER)

		//Act
		seller, err := sellerSv.Create(sellerAttributesB)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Seller{}, seller)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 4: foreign key - locality ID don't exist", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		errForeignKey := eh.GetErrForeignKey(eh.LOCALITY)

		localityRp.On("GetByID", 1).Return(model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY))

		//Act
		seller, err := sellerSv.Create(sellerAttributes)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, errForeignKey, err)
		require.Equal(t, model.Seller{}, seller)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 5: Invalid data - locality ID is not a number", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		errInvalidData := eh.GetErrInvalidData(eh.LOCALITY)

		sellerAttributes := model.SellerAttributes{
			LocalityId:  &badLocalityId,
		}

		localityRp.On("GetByID", 1).Return(model.LocalityDBModel{}, eh.GetErrInvalidData(eh.LOCALITY))

		//Act
		err := sellerSv.validateLocality(*sellerAttributes.LocalityId)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
	})
}

func TestSellerService_Update(t *testing.T) {
	t.Run("case 1: update seller successfully", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		seller := model.Seller{
			ID:               1,
			SellerAttributes: sellerAttributes,
		}

		localityRp.On("GetByID", 1).Return(model.LocalityDBModel{Id: 1}, nil)
		sellerRp.On("Update", seller.ID, sellerAttributes).Return(seller, nil)

		//Act
		sellerResult, err := sellerSv.Update(seller.ID, sellerAttributes)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, seller)
		require.Equal(t, seller, sellerResult)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 2: not found - seller not found", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		errNotFound := eh.GetErrNotFound(eh.SELLER)
		seller := model.Seller{
			ID:               1,
			SellerAttributes: sellerAttributes,
		}

		localityRp.On("GetByID", 1).Return(model.LocalityDBModel{Id: 1}, nil)
		sellerRp.On("Update", seller.ID, seller.SellerAttributes).Return(model.Seller{}, errNotFound)

		//Act
		sellerResult, err := sellerSv.Update(seller.ID, seller.SellerAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errNotFound, err)
		require.Equal(t, model.Seller{}, sellerResult)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 5: Unprocessable Entity - invalid data Company", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		errInvalidData := eh.GetErrInvalidData(eh.SELLER)
		seller := model.Seller{
			ID: 1,
			SellerAttributes: model.SellerAttributes{
				CompanyID:   &badCompanyID,
				CompanyName: &companyName,
				Address:     &address,
				Telephone:   &telephone,
				LocalityId:  &localityId,
			},
		}

		localityRp.On("GetByID", 1).Return(model.LocalityDBModel{Id: 1}, nil)
		sellerRp.On("Update", seller.ID, seller.SellerAttributes).Return(model.Seller{}, errInvalidData)

		//Act
		sellerResult, err := sellerSv.Update(seller.ID, seller.SellerAttributes)

		// Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrInvalidData)
		require.Equal(t, errInvalidData, err)
		require.Equal(t, model.Seller{}, sellerResult)
		sellerRp.AssertNotCalled(t, "Update")
		localityRp.AssertNotCalled(t, "Update")
	})
	t.Run("case 4: foreign key - locality ID don't exist", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		sellerSv := NewSellerService(sellerRp, localityRp)
		errForeignKey := eh.GetErrForeignKey(eh.LOCALITY)
		seller := model.Seller{
			ID:               1,
			SellerAttributes: sellerAttributes,
		}

		localityRp.On("GetByID", 1).Return(model.LocalityDBModel{}, eh.GetErrNotFound(eh.LOCALITY))

		//Act
		seller, err := sellerSv.Update(seller.ID, seller.SellerAttributes)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrForeignKey)
		require.Equal(t, errForeignKey, err)
		require.Equal(t, model.Seller{}, seller)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
}

func TestSellerService_Delete(t *testing.T) {
	t.Run("case 1: delete seller successfully", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()

		sellerSv := NewSellerService(sellerRp, localityRp)

		sellerRp.On("Delete", 1).Return(nil)

		//Act
		err := sellerSv.Delete(1)

		//Assert
		require.NoError(t, err)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
	t.Run("case 2:  not found - seller not found", func(t *testing.T) {
		//Arrange
		sellerRp := sellerRepository.NewSellerRepositoryMock()
		localityRp := sellerLocality.NewLocalityRepositoryMock()
		errNotFound := eh.GetErrNotFound(eh.SELLER)

		sellerSv := NewSellerService(sellerRp, localityRp)

		sellerRp.On("Delete", 2).Return(errNotFound)

		//Act
		err := sellerSv.Delete(2)

		//Assert
		require.Error(t, err)
		require.ErrorIs(t, err, eh.ErrNotFound)
		require.Equal(t, errNotFound, err)
		sellerRp.AssertExpectations(t)
		localityRp.AssertExpectations(t)
	})
}
