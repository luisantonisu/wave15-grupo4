package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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
		return nil, eh.GetErrNotFound(eh.SELLER)
	}

	return
}

func (s *SellerService) GetByID(id int) (model.Seller, error) {

	if id == 0 {
		return model.Seller{}, eh.GetErrNotFound(eh.SELLER)
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

func (s *SellerService) Update(id int, seller model.SellerAtrributesPtr) (model.Seller, error) {

	updatedSeller, err := s.rp.Update(id, seller)
	if err != nil {
		return model.Seller{}, err
	}

	return updatedSeller, nil

}

func (s *SellerService) Delete(id int) error {
	if id == 0 {
		return eh.GetErrNotFound(eh.SELLER)
	}
	return s.rp.Delete(id)
}

func (s *SellerService) validateSeller(seller model.Seller) error {

	hasCID := seller.CompanyID != 0
	hasCompanyName := seller.CompanyName != ""
	hasAddress := seller.Address != ""
	hasTelephone := seller.Telephone != ""

	if !hasCID || !hasCompanyName || !hasAddress || !hasTelephone {
		return eh.GetErrInvalidData(eh.SELLER)
	}

	return nil
}
