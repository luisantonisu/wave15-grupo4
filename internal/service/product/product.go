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

func (productService *ProductService) GetProductByID(id int) (product model.Product, err error) {
	return productService.repository.GetProductByID(id)
}

func ValueCheck(productAtrributes model.ProductAtrributes) (err error) {
	if productAtrributes.ProductCode == "" {
		return errors.New("ProductCode is invalid")
	}
	if productAtrributes.Description == "" {
		return errors.New("Description is invalid")
	}
	if productAtrributes.Width <= 0 {
		return errors.New("Width is invalid")
	}
	if productAtrributes.Height <= 0 {
		return errors.New("Height is invalid")
	}
	if productAtrributes.Length <= 0 {
		return errors.New("Length is invalid")
	}
	if productAtrributes.NetWeight <= 0 {
		return errors.New("NetWeight is invalid")
	}
	if productAtrributes.ExpirationRate <= 0 {
		return errors.New("ExpirationRate is invalid")
	}
	if productAtrributes.RecommendedFreezingTemperature <= 0 {
		return errors.New("RecommendedFreezingTemperature is invalid")
	}
	if productAtrributes.FreezingRate <= 0 {
		return errors.New("FreezingRate is invalid")
	}
	if productAtrributes.ProductTypeID <= 0 {
		return errors.New("ProductTypeId is invalid")
	}
	return
}

func (productService *ProductService) CreateProduct(productAttributes *model.ProductAtrributes) (err error) {
	if err = ValueCheck(*productAttributes); err != nil {
		return err
	}
	return productService.repository.CreateProduct(productAttributes)
}

func (productService *ProductService) DeleteProduct(id int) (err error) {
	return productService.repository.DeleteProduct(id)
}

func (productService *ProductService) UpdateProduct(id int, productAttributes *model.ProductAtrributes) (producto *model.Product, err error) {
	return productService.repository.UpdateProduct(id, productAttributes)
}
