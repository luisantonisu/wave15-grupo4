package dto

type SellerResponseDTO struct {
	ID          int    `json:"id"`
	CompanyID   string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellerRequestDTO struct {
	CompanyID   string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellerRequestDTOPtr struct {
	CompanyID   *string `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
}
