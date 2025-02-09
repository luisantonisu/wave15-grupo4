package dto

type ProductRecordRequestDTO struct {
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}

type ProductRecordResponseDTO struct {
	ID             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      int     `json:"product_id"`
}

type ProductRecordRequestDTOPtr struct {
	LastUpdateDate *string  `json:"last_update_date"`
	PurchasePrice  *float64 `json:"purchase_price"`
	SalePrice      *float64 `json:"sale_price"`
	ProductId      *int     `json:"product_id"`
}
