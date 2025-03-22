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

func (r *ProductRecordRepository) CreateProductRecord(productRecord model.ProductRecordAtrributes) (prodRecord model.ProductRecord, err error) {

	row, err := r.db.Exec("INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)", productRecord.LastUpdateDate, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductId)

	if err != nil {
		return model.ProductRecord{}, errorHandler.GetErrInvalidData(errorHandler.PRODUCT_RECORD)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.ProductRecord{}, err
	}

	prodRecord.ID = int(id)
	prodRecord.ProductRecordAtrributes = productRecord

	return
}
