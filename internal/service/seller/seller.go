package service

import (
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
)

func NewSellerService(rp repository.SellerRepository) *SellerService {
	return &SellerService{rp: rp}
}

type SellerService struct {
	rp repository.SellerRepository
}