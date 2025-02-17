package repository

import (
	"database/sql"
	"strconv"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewBuyerRepository(db *sql.DB) *BuyerRepository {
	return &BuyerRepository{
		db: db,
	}
}

type BuyerRepository struct {
	db *sql.DB
}

// Create a new buyer
func (r *BuyerRepository) Create(buyer model.BuyerAttributes) (model.Buyer, error) {
	// Validate card number id doesnt already exist
	if r.cardNumberIdExists(buyer.CardNumberId) {
		return model.Buyer{}, eh.GetErrAlreadyExists(eh.CARD_NUMBER)
	}
	// Create new buyer in DB
	row, err := r.db.Exec("INSERT INTO buyers (first_name, last_name, card_number_id) VALUES (?, ?, ?)",
		buyer.FirstName, buyer.LastName, buyer.CardNumberId,
	)
	if err != nil {
		return model.Buyer{}, eh.GetErrInvalidData(eh.BUYER)
	}
	id, err := row.LastInsertId()
	if err != nil {
		return model.Buyer{}, eh.GetErrDatabase(eh.BUYER)
	}

	// Response
	var newBuyer model.Buyer
	newBuyer.ID = int(id)
	newBuyer.BuyerAttributes = buyer

	return newBuyer, nil
}

// List all buyers
func (r *BuyerRepository) GetAll() ([]model.Buyer, error) {
	// Get buyers from db
	rows, err := r.db.Query("SELECT id, first_name, last_name, card_number_id FROM buyers")
	if err != nil {
		return nil, eh.GetErrGettingData(eh.BUYER)
	}
	defer rows.Close()

	// Parse buyers and create response
	var buyers []model.Buyer
	for rows.Next() {
		var buyer model.Buyer
		err := rows.Scan(&buyer.ID, &buyer.FirstName, &buyer.LastName, &buyer.CardNumberId)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.BUYER)
		}
		buyers = append(buyers, buyer)
	}
	return buyers, nil
}

// Get a buyer by id
func (r *BuyerRepository) GetByID(id int) (model.Buyer, error) {
	// Create response entity
	var buyer model.Buyer

	// Get buyer from db
	err := r.db.QueryRow("SELECT id, first_name, last_name, card_number_id FROM buyers WHERE id = ?", id).Scan(
		&buyer.ID, &buyer.FirstName, &buyer.LastName, &buyer.CardNumberId,
	)
	if err != nil {
		return model.Buyer{}, eh.GetErrNotFound(eh.BUYER)
	}
	return buyer, nil
}

// Delete a buyer by id
func (r *BuyerRepository) Delete(id int) error {
	// Verify buyer exists
	if !r.buyerExists(id) {
		return eh.GetErrNotFound(eh.BUYER)
	}

	// Delete buyer from db
	_, err := r.db.Exec("DELETE FROM buyers WHERE id = ?", id)
	if err != nil {
		return eh.GetErrDatabase(eh.BUYER)
	}
	return nil
}

// Update a buyer by id
func (r *BuyerRepository) Update(id int, buyer model.BuyerAttributesPtr) (model.Buyer, error) {
	// Validate buyer exists
	if !r.buyerExists(id) {
		return model.Buyer{}, eh.GetErrNotFound(eh.BUYER)
	}

	var newBuyer model.Buyer
	err := r.db.QueryRow("SELECT id, first_name, last_name, card_number_id FROM buyers WHERE id = ?", id).Scan(
		&newBuyer.ID, &newBuyer.FirstName, &newBuyer.LastName, &newBuyer.CardNumberId,
	)
	if err != nil {
		return model.Buyer{}, eh.GetErrDatabase(eh.BUYER)
	}

	// Update buyer entity with new values
	if buyer.FirstName != nil {
		newBuyer.FirstName = *buyer.FirstName
	}
	if buyer.LastName != nil {
		newBuyer.LastName = *buyer.LastName
	}
	if buyer.CardNumberId != nil {
		// Validate card number id already exist
		if r.cardNumberIdIsMine(*buyer.CardNumberId, id) {
			return model.Buyer{}, eh.GetErrAlreadyExists(eh.CARD_NUMBER)
		}
		newBuyer.CardNumberId = *buyer.CardNumberId
	}

	// Update buyer in db
	_, err = r.db.Exec("UPDATE buyers SET first_name = ?, last_name = ?, card_number_id = ? WHERE id = ?",
		newBuyer.FirstName, newBuyer.LastName, newBuyer.CardNumberId, id,
	)
	if err != nil {
		return model.Buyer{}, eh.GetErrInvalidData(eh.BUYER)
	}

	return newBuyer, nil
}

// Generate Purchase Order reports for buyer
func (r *BuyerRepository) PurchaseOrderReport(id int) ([]model.ReportPurchaseOrders, error) {
	// Make a "dynamic" query to get single or multiple buyers
	query := "SELECT buyers.id, buyers.first_name, buyers.last_name, buyers.card_number_id, COUNT(*) AS purchase_orders_count FROM buyers JOIN purchase_orders ON buyers.id = purchase_orders.buyer_id GROUP BY buyers.id"

	if id != -1 {
		// Verify buyer exists
		if !r.buyerExists(id) {
			return nil, eh.GetErrNotFound(eh.BUYER)
		}
		query += " HAVING buyers.id = " + strconv.Itoa(id)
	}
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, eh.GetErrGettingData(eh.BUYER)
	}

	// Response
	var buyers []model.ReportPurchaseOrders
	for rows.Next() {
		var buyer model.ReportPurchaseOrders
		err := rows.Scan(&buyer.ID, &buyer.FirstName, &buyer.LastName, &buyer.CardNumberId, &buyer.PurchaseOrdersCount)
		if err != nil {
			return nil, eh.GetErrParsingData(eh.BUYER)
		}
		buyers = append(buyers, buyer)
	}
	return buyers, nil

}

// Validate if card number id is already in use
func (r *BuyerRepository) cardNumberIdExists(cardNumberId string) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM buyers WHERE card_number_id = ?)", cardNumberId).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// Validate if card number is from the current user
func (r *BuyerRepository) cardNumberIdIsMine(cardNumberId string, id int) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM buyers WHERE card_number_id = ? AND id != ?)", cardNumberId, id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// Validate if a buyer exists
func (r *BuyerRepository) buyerExists(id int) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM buyers WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
