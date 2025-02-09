package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/product"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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

func (productService *ProductService) GetProductRecord() (productRecordMap map[int]model.ProductRecord, err error) {
	return productService.repository.GetProductRecord()
}

func (productService *ProductService) GetProductRecordByID(id int) (productRecord model.ProductRecord, err error) {
	return productService.repository.GetProductRecordByID(id)
}

func ValueCheck(productAtrributes model.ProductAtrributes) (err error) {
	if productAtrributes.ProductCode == "" || productAtrributes.Description == "" || productAtrributes.Width <= 0 || productAtrributes.Height <= 0 || productAtrributes.Length <= 0 || productAtrributes.NetWeight <= 0 || productAtrributes.ExpirationRate <= 0 || productAtrributes.RecommendedFreezingTemperature <= 0 || productAtrributes.FreezingRate <= 0 || productAtrributes.ProductTypeID <= 0 {
		return errorHandler.GetErrInvalidData(errorHandler.PRODUCT)
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

func (productService *ProductService) UpdateProduct(id int, productAttributes *model.ProductAtrributesPtr) (producto *model.Product, err error) {
	return productService.repository.UpdateProduct(id, productAttributes)
}
