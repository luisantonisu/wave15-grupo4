package dto

import eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"

type ProductBatchResponseDTO struct {
	ID                 int     `json:"id"`
	BatchNumber        string  `json:"batch_number"`
	CurrentQuantity    int     `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    int     `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  string  `json:"manufacturing_hour"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	ProductID          int     `json:"product_id"`
	SectionID          int     `json:"section_id"`
}

type ProductBatchRequestDTO struct {
	BatchNumber        string  `json:"batch_number"`
	CurrentQuantity    int     `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    int     `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  string  `json:"manufacturing_hour"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	ProductID          int     `json:"product_id"`
	SectionID          int     `json:"section_id"`
}

// Validate ProductBatchRequestDTO required fields
func (e *ProductBatchRequestDTO) Validate() error {
	if e.BatchNumber == "" {
		return eh.GetErrInvalidData(eh.PRODUCT_BATCH)
	}
	return nil
}
