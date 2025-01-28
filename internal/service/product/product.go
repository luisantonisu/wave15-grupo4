package service

import (
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
)

func NewProductService(rp repository.IProduct) *ProductService {
	return &ProductService{rp: rp}
}

type ProductService struct {
	rp repository.IProduct
}
