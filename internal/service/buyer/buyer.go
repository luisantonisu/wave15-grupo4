package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	buyerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
)

func NewBuyerService(rp buyerRepository.IBuyer) *BuyerService {
	return &BuyerService{rp: rp}
}

type BuyerService struct {
	rp buyerRepository.IBuyer
}

// Create a new buyer
func (s *BuyerService) Create(buyer model.BuyerAttributes) (model.Buyer, error) {
	newBuyer, err := s.rp.Create(buyer)
	if err != nil {
		return model.Buyer{}, err
	}
	return newBuyer, nil
}

// List all buyers
func (s *BuyerService) GetAll() ([]model.Buyer, error) {
	allBuyers, err := s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return allBuyers, nil
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
	err := s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// Update a buyer by id
func (s *BuyerService) Update(id int, buyer model.BuyerAttributes) (model.Buyer, error) {
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
