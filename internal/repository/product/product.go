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

func (productRepository *ProductRepository) GetProductById(id int) (product model.Product, err error) {
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
