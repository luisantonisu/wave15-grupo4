package model

type Product struct {
	ID int
	ProductAtrributes
}

type ProductRecordCount struct {
	ProductID   int
	Description string
	Count       int
}

type ProductAtrributes struct {
	ProductCode                    string
	Description                    string
	Width                          float64
	Height                         float64
	Length                         float64
	NetWeight                      float64
	ExpirationRate                 int
	RecommendedFreezingTemperature float64
	FreezingRate                   int
	ProductTypeID                  int
	SellerID                       int
}

type ProductAtrributesPtr struct {
	ProductCode                    *string
	Description                    *string
	Width                          *float64
	Height                         *float64
	Length                         *float64
	NetWeight                      *float64
	ExpirationRate                 *int
	RecommendedFreezingTemperature *float64
	FreezingRate                   *int
	ProductTypeID                  *int
	SellerID                       *int
}
