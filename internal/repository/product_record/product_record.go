package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	"github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type ProductRecordRepository struct {
	db *sql.DB
}

func NewProductRecordRepository(defaultDB *sql.DB) *ProductRecordRepository {
	return &ProductRecordRepository{
		db: defaultDB,
	}
}

func (productRecordRepository *ProductRecordRepository) GetProductRecord() (map[int]model.ProductRecord, error) {
	rows, err := productRecordRepository.db.Query("SELECT id, product_id, quantity, price FROM product_records")
	if err != nil {
		return nil, error_handler.GetErrNotFound(error_handler.PRODUCT_RECORD)
	}
	defer rows.Close()

	var productRecords = make(map[int]model.ProductRecord)
	for rows.Next() {
		var productRecord model.ProductRecord
		err := rows.Scan(&productRecord.ID, &productRecord.ProductRecordAtrributes.LastUpdateDate, &productRecord.ProductRecordAtrributes.PurchasePrice, &productRecord.ProductRecordAtrributes.SalePrice, &productRecord.ProductRecordAtrributes.ProductId)

		if err != nil {
			return nil, error_handler.GetErrNotFound(error_handler.PRODUCT_RECORD)
		}
		productRecords[productRecord.ID] = productRecord
	}

	return productRecords, nil
}

func (productRecordRepository *ProductRecordRepository) GetProductRecordByID(id int) (model.ProductRecord, error) {
	row := productRecordRepository.db.QueryRow("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records WHERE id = ?", id)
	var productRecord model.ProductRecord
	err := row.Scan(&productRecord.ID, &productRecord.ProductRecordAtrributes.LastUpdateDate, &productRecord.ProductRecordAtrributes.PurchasePrice, &productRecord.ProductRecordAtrributes.SalePrice, &productRecord.ProductRecordAtrributes.ProductId)
	if err != nil {
		return model.ProductRecord{}, error_handler.GetErrNotFound(error_handler.PRODUCT_RECORD)
	}
	return productRecord, nil
}

func (productRecordRepository *ProductRecordRepository) CreateProductRecord(productRecord model.ProductRecordAtrributes) error {

	_, err := productRecordRepository.db.Exec("INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)", productRecord.LastUpdateDate, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductId)

	if err != nil {
		return error_handler.GetErrInvalidData(error_handler.PRODUCT_RECORD)
	}

	return err
}

// func (productRepository *ProductRecordRepository) registerExists(id int) (bool, error) {
// 	var exist bool
// 	query := "SELECT EXISTS(SELECT 1 FROM product WHERE ID = ?)"
// 	err := productRepository.db.QueryRow(query, id).Scan(&exist)
// 	if err != nil {
// 		return false, errorHandler.GetErrNotFound(errorHandler.PRODUCT)
// 	}
// 	return exist, nil
// }
