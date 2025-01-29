package dto

type SellerDTO struct {
	Id          int    `json:"id"`
	CompanyId   int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}
