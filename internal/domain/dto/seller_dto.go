package dto

type SellerResponseDTO struct {
	ID          int    `json:"id"`
	CompanyID   int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellerRequestDTO struct {
	CompanyID   int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellerRequestDTOPtr struct {
	CompanyID   *int    `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
}
