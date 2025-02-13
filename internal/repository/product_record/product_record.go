package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type ProductRecordRepository struct {
	db *sql.DB
}

func NewProductRecordRepository(defaultDB *sql.DB) *ProductRecordRepository {
	return &ProductRecordRepository{
		db: defaultDB,
	}
}

// func (productRecordRepository *ProductRecordRepository) GetProductRecord() (map[int]model.ProductRecord, error) {
// 	rows, err := productRecordRepository.db.Query("SELECT id, product_id, quantity, price FROM product_records")
// 	if err != nil {
// 		return nil, errorHandler.GetErrNotFound(errorHandler.PRODUCT_RECORD)
// 	}
// 	defer rows.Close()

// 	var productRecords = make(map[int]model.ProductRecord)
// 	for rows.Next() {
// 		var productRecord model.ProductRecord
// 		err := rows.Scan(&productRecord.ID, &productRecord.ProductRecordAtrributes.LastUpdateDate, &productRecord.ProductRecordAtrributes.PurchasePrice, &productRecord.ProductRecordAtrributes.SalePrice, &productRecord.ProductRecordAtrributes.ProductId)

// 		if err != nil {
// 			return nil, errorHandler.GetErrNotFound(errorHandler.PRODUCT_RECORD)
// 		}
// 		productRecords[productRecord.ID] = productRecord
// 	}

// 	return productRecords, nil
// }

// func (productRecordRepository *ProductRecordRepository) GetProductRecordByID(id int) (model.ProductRecord, error) {
// 	row := productRecordRepository.db.QueryRow("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records WHERE id = ?", id)
// 	var productRecord model.ProductRecord
// 	err := row.Scan(&productRecord.ID, &productRecord.ProductRecordAtrributes.LastUpdateDate, &productRecord.ProductRecordAtrributes.PurchasePrice, &productRecord.ProductRecordAtrributes.SalePrice, &productRecord.ProductRecordAtrributes.ProductId)
// 	if err != nil {
// 		return model.ProductRecord{}, errorHandler.GetErrNotFound(errorHandler.PRODUCT_RECORD)
// 	}
// 	return productRecord, nil
// }

func (productRecordRepository *ProductRecordRepository) productIdExists(productId int) bool {
	row := productRecordRepository.db.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", productId)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (productRecordRepository *ProductRecordRepository) CreateProductRecord(productRecord model.ProductRecordAtrributes) error {

	productIdExists := productRecordRepository.productIdExists(productRecord.ProductId)

	if !productIdExists {
		return errorHandler.GetErrAlreadyExists("product doesn't")
	}

	_, err := productRecordRepository.db.Exec("INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)", productRecord.LastUpdateDate, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductId)

	if err != nil {
		return errorHandler.GetErrInvalidData("product record")
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
