package dto

import eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"

type BuyerResponseDTO struct {
	ID           int    `json:"id"`
	BuyerRequestDTO
}

type BuyerRequestDTO struct {
	CardNumberId *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
}

type ReportPurchaseOrdersResponseDTO struct {
	ID                  int    `json:"id"`
	CardNumberId        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

// Validate BuyerRequestDTO fields
func (e *BuyerRequestDTO) Validate() error {
	if e.CardNumberId == nil {
		return eh.GetErrInvalidData(eh.CARD_NUMBER)
	}
	return nil
}
