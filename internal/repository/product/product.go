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

func (pr *ProductRepository) GetProduct() (prMap map[int]model.Product, err error) {
	if len(pr.db) == 0 {
		return nil, errors.New("no products found")
	}
	return pr.db, nil
}

func (pr *ProductRepository) GetProductById(id int) (p model.Product, err error) {
	if len(pr.db) == 0 {
		return model.Product{}, errors.New("No products found")
	}
	for _, prod := range pr.db {
		if prod.ID == id {
			return prod, nil
		}
	}
	return model.Product{}, errors.New("Product not found")
}

func (pr *ProductRepository) CreateProduct(p *model.Product) (prod *model.Product, err error) {
	pr.db[len(pr.db)+1] = *p
	if p == nil {
		return nil, errors.New("Product is nil")
	}
	prod = p
	return p, nil
}
