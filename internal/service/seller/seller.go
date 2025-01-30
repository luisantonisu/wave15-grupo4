package service

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
)

func NewSellerService(rp repository.ISeller) *SellerService {
	return &SellerService{rp: rp}
}

type SellerService struct {
	rp repository.ISeller
}

func (s *SellerService) GetAll() (sellers map[int]model.Seller, err error) {
    sellers, err = s.rp.GetAll()
    if len(sellers) == 0 {
		return nil, errors.New("data not found")
	}

	return
}

func (s *SellerService) GetByID(id int) (model.Seller, error) {

	if id == 0 {
		return model.Seller{}, errors.New("data not found")
	}
	
	
	seller, err := s.rp.GetByID(id)
	if err != nil {
		return model.Seller{}, err
	}

	return seller, nil
}

func (s *SellerService) Create(seller model.Seller) (model.Seller, error) {
    
	//validate seller
	err := s.validateSeller(seller)
	if err != nil {
		return model.Seller{}, err
	}

	//create seller
	newSeller, err := s.rp.Create(seller)
	if err != nil {
		return model.Seller{}, err
	}

	return newSeller, nil
}

func (s *SellerService) validateSeller(seller model.Seller) error {
	
	hasCID := seller.CompanyID != 0
	hasCompanyName := seller.CompanyName != ""
	hasAddress := seller.Address != ""
	hasTelephone := seller.Telephone != ""

	if !hasCID || !hasCompanyName || !hasAddress || !hasTelephone {
		return errors.New("seller data incorrectly formed or incomplete")
	}

	return nil
}

