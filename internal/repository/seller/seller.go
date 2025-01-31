package repository

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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

func (r *SellerRepository) GetAll() (map[int]model.Seller, error) {
	sellers := make(map[int]model.Seller)

	for key, seller := range r.db {
		sellers[key] = seller
	}

	return sellers, nil
}

func (r *SellerRepository) GetByID(id int) (model.Seller, error) {

	for _, seller := range r.db {
		if seller.ID == id {
			return seller, nil
		}
	}

	return model.Seller{}, eh.GetErrNotFound(eh.SELLER)
}

func (r *SellerRepository) Create(seller model.Seller) (model.Seller, error) {
	//Create id
	id := len(r.db) + 1
	seller.ID = id

	//validate
	err := r.validateCompanyID(seller.CompanyID)
	if err != nil {
		return model.Seller{}, err
	}

	//Add new value
	r.db[id] = seller

	// return new seller
	return seller, nil
}

func (r *SellerRepository) Update(id int, seller model.SellerAtrributesPtr) (model.Seller, error) {
	if _, ok := r.db[id]; !ok {
		return model.Seller{}, eh.GetErrNotFound(eh.SELLER)
	}

	for _, value := range r.db {
		if value.CompanyID == *seller.CompanyID && value.ID != id {
			return model.Seller{}, eh.ErrAlreadyExists
		}
	}

	updateSeller := r.db[id]

	if seller.CompanyID != nil {
		updateSeller.CompanyID = *seller.CompanyID
	}

	if seller.CompanyName != nil {
		updateSeller.CompanyName = *seller.CompanyName
	}

	if seller.Address != nil {
		updateSeller.Address = *seller.Address
	}

	if seller.Telephone != nil {
		updateSeller.Telephone = *seller.Telephone
	}

	r.db[id] = updateSeller

	return updateSeller, nil

}

func (r *SellerRepository) Delete(id int) error {
	_, ok := r.db[id]
	if !ok {
		return eh.GetErrNotFound(eh.SELLER)
	}
	delete(r.db, id)
	return nil
}

func (r *SellerRepository) validateCompanyID(companyID int) error {
	for _, seller := range r.db {
		companyExist := seller.CompanyID == companyID

		if companyExist {
			return eh.GetErrAlreadyExists(eh.SELLER)
		}
	}

	return nil
}
