package service

import (
	"regexp"
	"strconv"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	localityRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/locality"
	sellerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewSellerService(sellerRp sellerRepository.ISeller, localityRp localityRepository.ILocality) *SellerService {
	return &SellerService{
		sellerRp:   sellerRp,
		localityRp: localityRp,
	}
}

type SellerService struct {
	sellerRp   sellerRepository.ISeller
	localityRp localityRepository.ILocality
}

func (s *SellerService) GetAll() (sellers []model.Seller, err error) {
	return s.sellerRp.GetAll()
}

func (s *SellerService) GetByID(id int) (model.Seller, error) {

	seller, err := s.sellerRp.GetByID(id)
	if err != nil {
		return model.Seller{}, err
	}

	return seller, nil
}

func (s *SellerService) Create(seller model.SellerAttributes) (model.Seller, error) {

	//validate seller
	err := s.ValidateSeller(seller)
	if err != nil {
		return model.Seller{}, err
	}

	//Validate if locality exist
	err = s.validateLocality(*seller.LocalityId)
	if err != nil {
		return model.Seller{}, eh.GetErrForeignKey(eh.LOCALITY)
	}

	//create seller
	newSeller, err := s.sellerRp.Create(seller)
	if err != nil {
		return model.Seller{}, err
	}

	return newSeller, nil
}

func (s *SellerService) Update(id int, seller model.SellerAttributes) (model.Seller, error) {
	if seller.LocalityId != nil {
		//Validate if locality exist
		err := s.validateLocality(*seller.LocalityId)
		if err != nil {
			return model.Seller{}, eh.GetErrForeignKey(eh.LOCALITY)
		}
	}

	pattern := regexp.MustCompile("^[1-9]+[0-9]*$")
	matchCompanyID := pattern.MatchString(*seller.CompanyID)
	if !matchCompanyID {
		return model.Seller{}, eh.GetErrInvalidData(eh.SELLER)
	}

	updatedSeller, err := s.sellerRp.Update(id, seller)
	if err != nil {
		return model.Seller{}, err
	}

	return updatedSeller, nil

}

func (s *SellerService) Delete(id int) error {
	return s.sellerRp.Delete(id)
}

func (s *SellerService) ValidateSeller(seller model.SellerAttributes) error {
	//validate if company_id only contains numbers and is not empty
	pattern := regexp.MustCompile("^[1-9]+[0-9]*$")
	matchCompanyID := pattern.MatchString(*seller.CompanyID)
	if !matchCompanyID {
		return eh.GetErrInvalidData(eh.SELLER)
	}

	//validate if locality_id only contains numbers and is not empty
	matchLocalityId := pattern.MatchString(*seller.LocalityId)
	if !matchLocalityId {
		return eh.GetErrInvalidData(eh.SELLER)
	}

	hasCompanyName := seller.CompanyName != nil && *seller.CompanyName != ""
	hasAddress := seller.Address != nil && *seller.Address != ""
	hasTelephone := seller.Telephone != nil && *seller.Telephone != ""

	var validators = []bool{hasCompanyName, hasAddress, hasTelephone}
	for _, item := range validators {
		if !item {
			return eh.GetErrInvalidData(eh.SELLER)
		}
	}

	return nil
}

func (s *SellerService) validateLocality(id string) error {
	//convert locality_id
	localityID, err := strconv.Atoi(id)
	if err != nil {
		return eh.GetErrInvalidData(eh.LOCALITY)
	}

	//Validate if locality exist
	_, err = s.localityRp.GetByID(localityID)
	if err != nil {
		return err
	}

	return nil
}
