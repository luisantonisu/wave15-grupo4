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

func (s *ProductService) GetProduct() (productMap map[int]model.Product, err error) {
	return s.repository.GetProduct()
}

func (s *ProductService) GetProductByID(id int) (product model.Product, err error) {
	return s.repository.GetProductByID(id)
}

func (s *ProductService) GetProductRecord() (productRecordMap map[int]model.ProductRecordCount, err error) {
	return s.repository.GetProductRecord()
}

func (s *ProductService) GetProductRecordByID(id int) (productRecord model.ProductRecordCount, err error) {
	return s.repository.GetProductRecordByID(id)
}

func ValueCheck(productAtrributes model.ProductAtrributes) (err error) {
	if productAtrributes.ProductCode == "" || productAtrributes.Description == "" || productAtrributes.Width <= 0 || productAtrributes.Height <= 0 || productAtrributes.Length <= 0 || productAtrributes.NetWeight <= 0 || productAtrributes.ExpirationRate <= 0 || productAtrributes.RecommendedFreezingTemperature <= 0 || productAtrributes.FreezingRate <= 0 || productAtrributes.ProductTypeID <= 0 {
		return errorHandler.GetErrInvalidData(errorHandler.PRODUCT)
	}
	return
}

func (s *ProductService) CreateProduct(productAttributes *model.ProductAtrributes) (prod model.Product, err error) {
	if err = ValueCheck(*productAttributes); err != nil {
		return model.Product{}, err
	}
	return s.repository.CreateProduct(productAttributes)
}

func (s *ProductService) DeleteProduct(id int) (err error) {
	return s.repository.DeleteProduct(id)
}

func (s *ProductService) UpdateProduct(id int, productAttributes *model.ProductAtrributesPtr) (producto *model.Product, err error) {
	return s.repository.UpdateProduct(id, productAttributes)
}
