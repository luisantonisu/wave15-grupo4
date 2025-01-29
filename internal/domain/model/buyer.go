package model

type Buyer struct {
	ID int
	BuyerAttributes
}

type BuyerAttributes struct {
	CardNumberId int
	FirstName    string
	LastName     string
}
