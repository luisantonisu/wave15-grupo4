package model

type Product struct {
	ID int
	ProductAtrributes
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
	ProductTypeId                  int
	SellerId                       int
}
