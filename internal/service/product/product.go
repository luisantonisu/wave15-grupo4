package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
)

func NewProductService(rp repository.IProduct) *ProductService {
	return &ProductService{rp: rp}
}

type ProductService struct {
	rp repository.IProduct
}

func (pr *ProductService) GetProduct() (prMap map[int]model.Product, err error) {
	return pr.rp.GetProduct()
}
