package model

type Product struct {
	ID int
	ProductAtrributes
}

type ProductAtrributes struct {
	ProductCode                    *string
	Description                    *string
	Width                          *float64
	Height                         *float64
	Length                         *float64
	NetWeight                      *float64
	ExpirationRate                 *float64
	RecommendedFreezingTemperature *float64
	FreezingRate                   *float64
	ProductTypeID                  *int
	SellerID                       *int
}

type ProductRecordCount struct {
	ProductID   int
	Description string
	Count       int
}
