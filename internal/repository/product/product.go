package repository

import (
	"errors"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func NewProductRepository(db map[int]model.Product) *ProductRepository {
	defaultDb := make(map[int]model.Product)
	if db != nil {
		defaultDb = db
	}
	return &ProductRepository{db: defaultDb}
}

type ProductRepository struct {
	db map[int]model.Product
}

func (productRepository *ProductRepository) GetProduct() (productMap map[int]model.Product, err error) {
	if len(productRepository.db) == 0 {
		return nil, errors.New("no products found")
	}
	return productRepository.db, nil
}

func (productRepository *ProductRepository) GetProductByID(id int) (product model.Product, err error) {
	if len(productRepository.db) == 0 {
		return model.Product{}, errors.New("No products found")
	}
	for _, prod := range productRepository.db {
		if prod.ID == id {
			return prod, nil
		}
	}
	return model.Product{}, errors.New("Product not found")
}

func (productRepository *ProductRepository) CreateProduct(productAtrributes *model.ProductAtrributes) (err error) {
	for _, prod := range productRepository.db {
		if prod.ProductAtrributes.ProductCode == productAtrributes.ProductCode {
			return errors.New("Product already exists")
		}
	}
	var newProduct model.Product
	newProduct.ID = len(productRepository.db) + 1
	newProduct.ProductAtrributes = *productAtrributes
	if productAtrributes == nil {
		return errors.New("Product is nil")
	}
	productRepository.db[len(productRepository.db)+1] = newProduct
	return nil
}

func (productRepository *ProductRepository) DeleteProduct(id int) (err error) {
	_, ok := productRepository.db[id]
	if !ok {
		return errors.New("Product not found")
	}
	delete(productRepository.db, id)
	return nil
}

func (productRepository *ProductRepository) UpdateProduct(id int, productAtrributes *model.ProductAtrributes) (product *model.Product, err error) {

	if _, ok := productRepository.db[id]; !ok {
		return nil, errors.New("Product not found")
	}

	if productAtrributes == nil {
		return nil, errors.New("Product is nil")
	}
	for _, prod := range productRepository.db {
		if prod.ProductAtrributes.ProductCode == productAtrributes.ProductCode {
			return nil, errors.New("Product code already exists")
		}
	}
	patchedProduct := productRepository.db[id]

	if productAtrributes.ProductCode != "" {
		patchedProduct.ProductAtrributes.ProductCode = productAtrributes.ProductCode
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.Description != "" {
		patchedProduct.ProductAtrributes.Description = productAtrributes.Description
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.Width != 0 {
		patchedProduct.ProductAtrributes.Width = productAtrributes.Width
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.Height != 0 {
		patchedProduct.ProductAtrributes.Height = productAtrributes.Height
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.Length != 0 {
		patchedProduct.ProductAtrributes.Length = productAtrributes.Length
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.NetWeight != 0 {
		patchedProduct.ProductAtrributes.NetWeight = productAtrributes.NetWeight
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.ExpirationRate != 0 {
		patchedProduct.ProductAtrributes.ExpirationRate = productAtrributes.ExpirationRate
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.RecommendedFreezingTemperature != 0 {
		patchedProduct.ProductAtrributes.RecommendedFreezingTemperature = productAtrributes.RecommendedFreezingTemperature
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.FreezingRate != 0 {
		patchedProduct.ProductAtrributes.FreezingRate = productAtrributes.FreezingRate
		productRepository.db[id] = patchedProduct
	}
	if productAtrributes.ProductTypeID != 0 {
		patchedProduct.ProductAtrributes.ProductTypeID = productAtrributes.ProductTypeID
		productRepository.db[id] = patchedProduct
	}
	patchedProduct = productRepository.db[id]
	product = &patchedProduct
	return product, nil
}
