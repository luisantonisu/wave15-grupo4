package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewSellerRepository(db *sql.DB) *SellerRepository {

	return &SellerRepository{db: db}
}

type SellerRepository struct {
	db *sql.DB
}

func (r *SellerRepository) GetAll() (map[int]model.Seller, error) {
	rows, err := r.db.Query("SELECT id, company_id, company_name, address, telephone FROM sellers")
	if err != nil {
		return nil, eh.GetErrNotFound(eh.SELLER)
	}

	sellersList := make(map[int]model.Seller)
	for rows.Next() {
		var seller model.Seller
		err := rows.Scan(&seller.ID, &seller.CompanyID, &seller.CompanyName, &seller.Address, &seller.Telephone)
		if err != nil {
			return nil, eh.GetErrNotFound(eh.SELLER)
		}

		sellersList[seller.ID] = seller
	}

	return sellersList, nil
}

func (r *SellerRepository) GetByID(id int) (model.Seller, error) {
	var seller model.Seller
	err := r.db.QueryRow("SELECT id, company_id, company_name, address, telephone FROM sellers WHERE id = ?", id).Scan(
		&seller.ID, &seller.CompanyID, &seller.CompanyName, &seller.Address, &seller.Telephone)
	if err != nil {
		return model.Seller{}, eh.GetErrNotFound(eh.SELLER)
	}

	return seller, nil
}

func (r *SellerRepository) Create(seller model.SellerAttributes) (model.Seller, error) {

	hasIdAlreadyExist := r.CompanyIDExist(seller.CompanyID)
	if hasIdAlreadyExist {
		return model.Seller{}, eh.GetErrAlreadyExists(eh.SELLER)
	}

	row, err := r.db.Exec("INSERT INTO sellers (company_id, company_name, address, telephone) VALUES (?,?,?,?)",
		seller.CompanyID, seller.CompanyName, seller.Address, seller.Telephone)
	if err != nil {
		return model.Seller{}, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.Seller{}, err
	}

	var newSeller model.Seller
	newSeller.ID = int(id)
	newSeller.SellerAttributes = seller

	// return new seller
	return newSeller, nil
}

func (r *SellerRepository) Update(id int, seller model.SellerAttributesPtr) (model.Seller, error) {
	//Verify if seller exist
	var updateSeller model.Seller
	err := r.db.QueryRow("SELECT id, company_id, company_name, address, telephone FROM sellers WHERE id = ?", id).Scan(
		&updateSeller.ID, &updateSeller.CompanyID, &updateSeller.CompanyName, &updateSeller.Address, &updateSeller.Telephone)
	if err != nil {
		return model.Seller{}, eh.GetErrNotFound(eh.SELLER)
	}

	if seller.CompanyID != nil {

		if r.companyIDBelongToSeller(*seller.CompanyID, id) {
			return model.Seller{}, eh.GetErrAlreadyExists(eh.COMPANY_ID)
		}

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

	_, err = r.db.Exec("UPDATE sellers SET company_id = ?, company_name = ?, address = ?, telephone = ? WHERE id = ?",
		updateSeller.CompanyID, updateSeller.CompanyName, updateSeller.Address, updateSeller.Telephone, id)
	if err != nil {
		return model.Seller{}, eh.GetErrInvalidData(eh.SELLER)
	}

	return updateSeller, nil

}

func (r *SellerRepository) Delete(id int) error {
	if !r.sellerExist(id) {
		return eh.GetErrNotFound(eh.SELLER)
	}

	_, err := r.db.Exec("DELETE FROM sellers WHERE id = ?", id)
	if err != nil {
		return eh.GetErrNotFound(eh.SELLER) 
	}

	return nil
}

func (r *SellerRepository) CompanyIDExist(companyID string) bool {
	var exist bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sellers WHERE ID = ?)", companyID).Scan(&exist)
	if err != nil {
		return false
	}

	return exist

}

func (r *SellerRepository) sellerExist(id int) bool {
	var exist bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sellers WHERE ID = ?)", id).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}

func (r *SellerRepository) companyIDBelongToSeller(companyId string, id int) bool {
	var exist bool
	//If seller exist and don't belong to seller
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM sellers WHERE company_id = ?) AND id != ?", companyId, id).Scan(&exist)
	if err != nil {
		return false
	}

	return exist
}
