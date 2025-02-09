package model

type ProductRecord struct {
	ID int
	ProductRecordAtrributes
}

type ProductRecordAtrributes struct {
	LastUpdateDate string
	PurchasePrice  string
	SalePrice      float64
	ProductId      float64
}

type ProductRecordAtrributesPtr struct {
	LastUpdateDate *string
	PurchasePrice  *string
	SalePrice      *float64
	ProductId      *float64
}
