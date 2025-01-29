package service

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
)

func NewProductService(repository repository.IProduct) *ProductService {
	return &ProductService{repository: repository}
}

type ProductService struct {
	repository repository.IProduct
}

func (productService *ProductService) GetProduct() (productMap map[int]model.Product, err error) {
	return productService.repository.GetProduct()
}

func (productService *ProductService) GetProductById(id int) (product model.Product, err error) {
	return productService.repository.GetProductById(id)
}

func (productService *ProductService) CreateProduct(product *model.ProductAtrributes) (err error) {

	if product.ProductCode == "" {
		return errors.New("ProductCode is invalid")
	}
	if product.Description == "" {
		return errors.New("Description is invalid")
	}
	if product.Width <= 0 {
		return errors.New("Width is invalid")
	}
	if product.Height <= 0 {
		return errors.New("Height is invalid")
	}
	if product.Length <= 0 {
		return errors.New("Length is invalid")
	}
	if product.NetWeight <= 0 {
		return errors.New("NetWeight is invalid")
	}
	if product.ExpirationRate <= 0 {
		return errors.New("ExpirationRate is invalid")
	}
	if product.RecommendedFreezingTemperature <= 0 {
		return errors.New("RecommendedFreezingTemperature is invalid")
	}
	if product.FreezingRate <= 0 {
		return errors.New("FreezingRate is invalid")
	}
	if product.ProductTypeId <= 0 {
		return errors.New("ProductTypeId is invalid")
	}
	return productService.repository.CreateProduct(product)
}

func (productService *ProductService) DeleteProduct(id int) (err error) {
	return productService.repository.DeleteProduct(id)
}
