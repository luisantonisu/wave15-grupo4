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

func (productRecordRepository *ProductRecordRepository) CreateProductRecord(productRecord model.ProductRecordAtrributes) error {

	_, err := productRecordRepository.db.Exec("INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)", productRecord.LastUpdateDate, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductId)

	if err != nil {
		return errorHandler.GetErrInvalidData(errorHandler.PRODUCT_RECORD)
	}

	return err
}
