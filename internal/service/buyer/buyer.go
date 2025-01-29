package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
)

func NewBuyerService(rp repository.IBuyer) *BuyerService {
	return &BuyerService{rp: rp}
}

type BuyerService struct {
	rp repository.IBuyer
}

// Create a new buyer
func (s *BuyerService) Create(buyer model.Buyer) (model.Buyer, error) {
	newBuyer, err := s.rp.Create(buyer)
	if err != nil {
		return model.Buyer{}, err
	}
	return newBuyer, nil
}

// List all buyers
func (s *BuyerService) GetAll() (map[int]model.Buyer, error) {
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
