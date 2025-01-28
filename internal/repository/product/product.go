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
