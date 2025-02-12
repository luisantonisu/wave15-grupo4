package model

type Buyer struct {
	ID int
	BuyerAttributes
}

type BuyerAttributes struct {
	CardNumberId string
	FirstName    string
	LastName     string
}

type BuyerAttributesPtr struct {
	CardNumberId *string
	FirstName    *string
	LastName     *string
}
