package model

type Seller struct {
	ID int
	SellerAttributes
}

type SellerAttributes struct {
	CompanyID   string
	CompanyName string
	Address     string
	Telephone   string
	LocalityId  string
}

type SellerAttributesPtr struct {
	CompanyID   *string
	CompanyName *string
	Address     *string
	Telephone   *string
	LocalityId  *string
}
