package model

import "fmt"

type Buyer struct {
	Id int
	BuyerAttributes
}

type BuyerAttributes struct {
	CardNumberId int
	FirstName    string
	LastName     string
}

// Validar todos los campos del buyer
func (e Buyer) Validate() error {
	if e.Id == 0 {
		return fmt.Errorf("Id is required")
	}
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
