package dto

type SellerResponseDTO struct {
	ID          int    `json:"id"`
	SellerRequestDTO
}

type SellerRequestDTO struct {
	CompanyID   *string `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
	LocalityId  *string `json:"locality_id"`
}
