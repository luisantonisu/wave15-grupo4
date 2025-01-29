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
