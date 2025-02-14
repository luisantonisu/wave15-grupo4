package dto

type CarryResponseDTO struct {
	ID          int    `json:"id"`
	CarryID     string `json:"carry_id"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   uint   `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

type CarryRequestDTO struct {
	CarryID     string `json:"carry_id"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   uint   `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

type CarryRequestDTOPtr struct {
	CarryID     *string `json:"carry_id"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *uint   `json:"telephone"`
	LocalityID  *int    `json:"locality_id"`
}
