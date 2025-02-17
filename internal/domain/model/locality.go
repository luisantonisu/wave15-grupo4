package model

type Locality struct {
	Id string
	LocalityAttributes
}

type LocalityAttributes struct {
	LocalityName string
	ProvinceName string
	CountryName  string
}

type LocalityAttributesPtr struct {
	LocalityName *string
	ProvinceName *string
	CountryName  *string
}

type LocalityDBModel struct {
	Id           int
	LocalityName string
	ProvinceID   int
}

type LocalityReport struct {
	Id           int 
	LocalityName string 
	SellerCount  int    
}
