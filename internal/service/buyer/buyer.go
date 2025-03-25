package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	buyerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
	"errors"
)

func NewBuyerService(rp buyerRepository.IBuyer) *BuyerService {
	return &BuyerService{rp: rp}
}

type BuyerService struct {
	rp buyerRepository.IBuyer
}

// Create a new buyer
func (s *BuyerService) Create(buyer model.BuyerAttributes) (model.Buyer, error) {
	// Validate card number id doesnt already exist
	exists, err := s.rp.GetByCardNumberID(*buyer.CardNumberId)
	if exists != (model.Buyer{}) {
		return model.Buyer{}, eh.GetErrAlreadyExists(eh.CARD_NUMBER) 
	}
	if err != nil && !errors.Is(err, eh.ErrNotFound) {
		return model.Buyer{}, err
	}

	// Create new buyer
	newBuyer, err := s.rp.Create(buyer)
	if err != nil {
		return model.Buyer{}, err
	}
	return newBuyer, nil
}

// List all buyers
func (s *BuyerService) GetAll() ([]model.Buyer, error) {
	return s.rp.GetAll()
}

// Get a buyer by id
func (s *BuyerService) GetByID(id int) (model.Buyer, error) {
	buyer, err := s.rp.GetByID(id)
	if err != nil {
		return model.Buyer{}, err
	}
	return buyer, nil
}

// Delete a buyer by id
func (s *BuyerService) Delete(id int) error {
	// Validate buyer exists
	_, err := s.rp.GetByID(id)
	if err != nil {
		return err
	}
	
	// Delete the buyer if exists
	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update a buyer by id
func (s *BuyerService) Update(id int, buyer model.BuyerAttributes) (model.Buyer, error) {
	// Validate buyer exists
	existingBuyer, err := s.rp.GetByID(id)
	if err != nil {
		return model.Buyer{}, err
	}

	// Validate new card number id doesn't already exist or it's the same
	exists, err := s.rp.GetByCardNumberID(*buyer.CardNumberId)
	if exists != (model.Buyer{}) && exists.ID != existingBuyer.ID {
		return model.Buyer{}, eh.GetErrAlreadyExists(eh.CARD_NUMBER)	
	}
	if err != nil && !errors.Is(err, eh.ErrNotFound) {
		return model.Buyer{}, err
	}
	// Update buyer
	updatedBuyer, err := s.rp.Update(id, buyer)
	if err != nil {
		return model.Buyer{}, err
	}
	return updatedBuyer, nil
}

// Generate Purchase Order Report
func (s *BuyerService) PurchaseOrderReport(id *int) ([]model.ReportPurchaseOrders, error) {
	report, err := s.rp.PurchaseOrderReport(id)
	if err != nil {
		return nil, err
	}
	return report, nil
}
