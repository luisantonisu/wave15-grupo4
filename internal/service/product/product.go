package service

import (
	"errors"

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

func (pr *ProductService) GetProductById(id int) (p model.Product, err error) {
	return pr.rp.GetProductById(id)
}

func (pr *ProductService) CreateProduct(p *model.Product) (prod *model.Product, err error) {

	if p.ProductCode == "" {
		return nil, errors.New("ProductCode is invalid")
	}
	if p.Description == "" {
		return nil, errors.New("Description is invalid")
	}
	if p.Width <= 0 {
		return nil, errors.New("Width is invalid")
	}
	if p.Height <= 0 {
		return nil, errors.New("Height is invalid")
	}
	if p.Length <= 0 {
		return nil, errors.New("Length is invalid")
	}
	if p.NetWeight <= 0 {
		return nil, errors.New("NetWeight is invalid")
	}
	if p.ExpirationRate <= 0 {
		return nil, errors.New("ExpirationRate is invalid")
	}
	if p.RecommendedFreezingTemperature <= 0 {
		return nil, errors.New("RecommendedFreezingTemperature is invalid")
	}
	if p.FreezingRate <= 0 {
		return nil, errors.New("FreezingRate is invalid")
	}
	if p.ProductTypeId <= 0 {
		return nil, errors.New("ProductTypeId is invalid")
	}
	return pr.rp.CreateProduct(p)
}
