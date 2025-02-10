package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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
	buyer.ID = r.getNextId()
	// Validate card number id doesnt already exist
	if err := r.validateCardNumberId(buyer.CardNumberId); err != nil {
		return model.Buyer{}, eh.GetErrAlreadyExists(eh.CARD_NUMBER)
	}
	// Create buyer in db
	r.db[buyer.ID] = buyer
	return buyer, nil
}

// List all buyers
func (r *BuyerRepository) GetAll() (map[int]model.Buyer, error) {
	return r.db, nil
}

// Get a buyer by id
func (r *BuyerRepository) GetByID(id int) (model.Buyer, error) {
	buyer, ok := r.db[id]
	if !ok {
		return model.Buyer{}, eh.GetErrNotFound(eh.BUYER)
	}
	return buyer, nil
}

// Delete a buyer by id
func (r *BuyerRepository) Delete(id int) error {
	_, ok := r.db[id]
	if !ok {
		return eh.GetErrNotFound(eh.BUYER)
	}
	delete(r.db, id)
	return nil
}

// Update a buyer by id
func (r *BuyerRepository) Update(id int, buyer model.BuyerAttributesPtr) (model.Buyer, error) {
	// Validate buyer exists
	buyerOld, ok := r.db[id]
	if !ok {
		return model.Buyer{}, eh.GetErrNotFound(eh.BUYER)
	}

	// Update buyer in db
	if buyer.FirstName != nil {
		buyerOld.FirstName = *buyer.FirstName
	}
	if buyer.LastName != nil {
		buyerOld.LastName = *buyer.LastName
	}
	if buyer.CardNumberId != nil {
		// Validate card number id doesnt already exist
		if err := r.validateCardNumberId(*buyer.CardNumberId); err != nil {
			return model.Buyer{}, eh.GetErrAlreadyExists(eh.CARD_NUMBER)
		}
		buyerOld.CardNumberId = *buyer.CardNumberId
	}

	r.db[id] = buyerOld
	return buyerOld, nil
}

// Get the nex available id for a new buyer
func (r *BuyerRepository) getNextId() int {
	newId := len(r.db) + 1
	return newId
}

// Validate if card number id is already in use
func (r *BuyerRepository) validateCardNumberId(cardNumberId int) error {
	for _, buyer := range r.db {
		if buyer.CardNumberId == cardNumberId {
			return eh.ErrAlreadyExists
		}
	}
	return nil
}
