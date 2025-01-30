package dto

import eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"

type BuyerRequestDTO struct {
	CardNumberId int    `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerResponseDTO struct {
	ID           int    `json:"id"`
	CardNumberId int    `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerRequestDTOPtr struct {
	CardNumberId *int    `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
}

// Validate BuyerRequestDTO fields
func (e *BuyerRequestDTO) Validate() error {
	if e.CardNumberId == 0 {
		return eh.GetErrInvalidData(eh.CARD_NUMBER)
	}
	return nil
}
