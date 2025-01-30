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

func (productRepository *ProductRepository) UpdateProduct(id int, productAtrributesPtr *model.ProductAtrributesPtr) (product *model.Product, err error) {

	if _, ok := productRepository.db[id]; !ok {
		return nil, errors.New("Product not found")
	}

	if productAtrributesPtr == nil {
		return nil, errors.New("Product is nil")
	}

	for _, prod := range productRepository.db {
		if prod.ProductAtrributes.ProductCode == *productAtrributesPtr.ProductCode {
			return nil, errors.New("Product code already exists")
		}
	}
	patchedProduct := productRepository.db[id]

	if productAtrributesPtr.ProductCode != nil {
		patchedProduct.ProductAtrributes.ProductCode = *productAtrributesPtr.ProductCode
	}
	if productAtrributesPtr.Description != nil {
		patchedProduct.ProductAtrributes.Description = *productAtrributesPtr.Description
	}
	if productAtrributesPtr.Width != nil {
		patchedProduct.ProductAtrributes.Width = *productAtrributesPtr.Width
	}
	if productAtrributesPtr.Height != nil {
		patchedProduct.ProductAtrributes.Height = *productAtrributesPtr.Height
	}
	if productAtrributesPtr.Length != nil {
		patchedProduct.ProductAtrributes.Length = *productAtrributesPtr.Length
	}
	if productAtrributesPtr.NetWeight != nil {
		patchedProduct.ProductAtrributes.NetWeight = *productAtrributesPtr.NetWeight
	}
	if productAtrributesPtr.ExpirationRate != nil {
		patchedProduct.ProductAtrributes.ExpirationRate = *productAtrributesPtr.ExpirationRate
	}
	if productAtrributesPtr.RecommendedFreezingTemperature != nil {
		patchedProduct.ProductAtrributes.RecommendedFreezingTemperature = *productAtrributesPtr.RecommendedFreezingTemperature
	}
	if productAtrributesPtr.FreezingRate != nil {
		patchedProduct.ProductAtrributes.FreezingRate = *productAtrributesPtr.FreezingRate
	}
	if productAtrributesPtr.ProductTypeID != nil {
		patchedProduct.ProductAtrributes.ProductTypeID = *productAtrributesPtr.ProductTypeID
	}

	// Update the product in the repository after all fields have been patched
	productRepository.db[id] = patchedProduct
	product = &patchedProduct
	return product, nil
}
