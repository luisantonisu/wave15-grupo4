package dto

type ProductRecordCountResponseDTO struct {
	ProductID   int    `json:"product_id"`
	Description string `json:"description"`
	Count       int    `json:"records_count"`
}

type ProductRequestDTO struct {
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ExpirationRate                 int     `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   int     `json:"freezing_rate"`
	ProductTypeId                  int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

type ProductResponseDTO struct {
	ID                             int     `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ExpirationRate                 int     `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   int     `json:"freezing_rate"`
	ProductTypeId                  int     `json:"product_type_id"`
	SellerId                       int     `json:"seller_id"`
}

type ProductRequestDTOPtr struct {
	ProductCode                    *string  `json:"product_code"`
	Description                    *string  `json:"description"`
	Width                          *float64 `json:"width"`
	Height                         *float64 `json:"height"`
	Length                         *float64 `json:"length"`
	NetWeight                      *float64 `json:"net_weight"`
	ExpirationRate                 *int     `json:"expiration_rate"`
	RecommendedFreezingTemperature *float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   *int     `json:"freezing_rate"`
	ProductTypeId                  *int     `json:"product_type_id"`
	SellerId                       *int     `json:"seller_id"`
}
