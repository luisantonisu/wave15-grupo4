package model

type Carry struct {
	ID int
	CarryAttributes
}

type CarryAttributes struct {
	CarryID     *string
	CompanyName *string
	Address     *string
	Telephone   *uint
	LocalityID  *int
}
