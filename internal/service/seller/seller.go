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
	seller, err := s.rp.GetByID(id)
	if err != nil {
		return model.Seller{}, err
	}

	return seller, nil
}