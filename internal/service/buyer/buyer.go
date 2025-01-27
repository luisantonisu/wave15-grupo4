package service

import (
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
)

func NewBuyerService(rp repository.BuyerRepository) *BuyerService {
	return &BuyerService{rp: rp}
}

type BuyerService struct {
	rp repository.BuyerRepository
}
