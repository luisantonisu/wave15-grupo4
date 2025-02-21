package repsotory

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewProductBatchRepository(db *sql.DB) *ProductBatchRepository {
	return &ProductBatchRepository{
		db: db,
	}
}

type ProductBatchRepository struct {
	db *sql.DB
}

func (p *ProductBatchRepository) productBatchExists(BatchNumber string) bool {
	var exists bool
	err := p.db.QueryRow("SELECT EXISTS(SELECT 1 FROM product_batches WHERE batch_number = ?)", BatchNumber).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (p *ProductBatchRepository) Create(productBatch model.ProductBatchAttributes) (model.ProductBatch, error) {
	if p.productBatchExists(productBatch.BatchNumber) {
		return model.ProductBatch{}, eh.GetErrAlreadyExists(eh.PRODUCT_BATCH_ID)
	}

	row, err := p.db.Exec("INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, intial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		productBatch.BatchNumber, productBatch.CurrentQuantity, productBatch.CurrentTemperature, productBatch.DueDate, productBatch.InitialQuantity, productBatch.ManufacturingDate, productBatch.ManufacturingHour, productBatch.MinimumTemperature, productBatch.ProductID, productBatch.SectionID)
	if err != nil {
		return model.ProductBatch{}, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return model.ProductBatch{}, err
	}

	var pBatch model.ProductBatch
	pBatch.ID = int(id)
	pBatch.ProductBatchAttributes = productBatch

	return pBatch, nil

}
