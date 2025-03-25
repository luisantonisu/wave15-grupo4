package dto

type CarryResponseDTO struct {
	ID int `json:"id"`
	CarryRequestDTO
}

type CarryRequestDTO struct {
	CarryID     *string `json:"carry_id"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
	LocalityID  *int    `json:"locality_id"`
}
