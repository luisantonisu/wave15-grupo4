package service

import (
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/seller"
)

func NewSellerService(rp repository.ISeller) *SellerService {
	return &SellerService{rp: rp}
}

type SellerService struct {
	rp repository.ISeller
}