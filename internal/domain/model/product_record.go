package model

type ProductRecord struct {
	ID int
	ProductRecordAtrributes
}

type ProductRecordAtrributes struct {
	LastUpdateDate string
	PurchasePrice  float64
	SalePrice      float64
	ProductId      int
}

type ProductRecordAtrributesPtr struct {
	LastUpdateDate *string
	PurchasePrice  *float64
	SalePrice      *float64
	ProductId      *int
}
