package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func NewSellerRepository(db map[int]model.Seller) *SellerRepository {
	defaultDb := make(map[int]model.Seller)
	if db != nil {
		defaultDb = db
	}
	return &SellerRepository{db: defaultDb}
}

type SellerRepository struct {
	db map[int]model.Seller
}

func (s *SellerRepository) GetAll() (sellers map[int]model.Seller, err error) {
	sellers = make(map[int]model.Seller)

	for key, seller := range s.db {
		sellers[key] = seller
	}
	
	return
}
