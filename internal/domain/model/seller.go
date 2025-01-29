package model

type Seller struct {
	ID int
	SellerAtrributes
}

type SellerAtrributes struct {
	CompanyID   int
	CompanyName string
	Address     string
	Telephone   string
}
