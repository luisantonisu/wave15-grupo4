package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func NewBuyerRepository(db map[int]model.Buyer) *BuyerRepository {
	defaultDb := make(map[int]model.Buyer)
	if db != nil {
		defaultDb = db
	}
	return &BuyerRepository{db: defaultDb}
}

type BuyerRepository struct {
	db map[int]model.Buyer
}

// Create a new buyer
func (r *BuyerRepository) Create(buyer model.Buyer) (model.Buyer, error) {
	buyer.Id = r.getNextId()
	r.db[buyer.Id] = buyer
	return buyer, nil
}

// Get the nex available id for a new buyer
func (r *BuyerRepository) getNextId() int {
	newId := len(r.db) + 1
	return newId
}
