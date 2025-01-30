package dto

import "fmt"

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
		return fmt.Errorf("CardNumberId is required")
	}
	if e.FirstName == "" {
		return fmt.Errorf("FirstName is required")
	}
	if e.LastName == "" {
		return fmt.Errorf("LastName is required")
	}
	return nil
}
