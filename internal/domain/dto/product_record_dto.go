package dto

type ProductRecordRequestDTO struct {
	LastUpdateDate string  `json:"last_update_code"`
	PurchasePrice  string  `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      float64 `json:"product_id"`
}

type ProductRecordResponseDTO struct {
	ID             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_code"`
	PurchasePrice  string  `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductId      float64 `json:"product_id"`
}

type ProductRecordRequestDTOPtr struct {
	LastUpdateDate *string  `json:"last_update_code"`
	PurchasePrice  *string  `json:"purchase_price"`
	SalePrice      *float64 `json:"sale_price"`
	ProductId      *float64 `json:"product_id"`
}
