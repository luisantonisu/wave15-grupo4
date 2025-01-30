package repository

import (
	"errors"

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

	return model.Seller{}, errors.New("seller not found")
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

func (r *SellerRepository) Update(id int, seller model.Seller) (model.Seller, error) {
	if _, ok := r .db[id]; !ok {
		return model.Seller{}, errors.New("seller not found")
	}

	for _, value := range r.db {
		if value.CompanyID == seller.CompanyID {
			return model.Seller{}, errors.New("a seller is already registered with this Company id")
		}
	}

	updateSeller := r.db[id]

	if seller.CompanyID > 0 {
		updateSeller.CompanyID = seller.CompanyID
	}

	if seller.CompanyName != "" {
		updateSeller.CompanyName = seller.CompanyName
	}

	if seller.Address != "" {
		updateSeller.Address = seller.Address 
	}

	if seller.Telephone != "" {
		updateSeller.Telephone = seller.Telephone
	}

	r.db[id] = updateSeller

	return updateSeller, nil

}

func (r *SellerRepository) validateCompanyID(companyID int) error {
	for _, seller := range r.db {
		companyExist := seller.CompanyID == companyID

		if companyExist {
			return errors.New("seller alredy exist")
		}
	}

	return nil
}
