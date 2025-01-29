package model

type Seller struct {
	Id int
	SellerAtrributes
}

type SellerAtrributes struct {
	CompanyId   int
	CompanyName string
	Address     string
	Telephone   string
}
