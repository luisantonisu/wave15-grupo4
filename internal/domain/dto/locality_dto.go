package dto

type LocalityRequestDTO struct {
	Data LocalityDataRequestDTO `json:"data"`
}

type LocalityDataRequestDTO struct {
	Id           string `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type LocalityDataResponseDTO struct {
	Id           string `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type LocalityResponseDTO struct {
	Id           string `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellerCount  int    `json:"sellers_count"`
}

type LocalityRequestDTOPtr struct {
	Id           *string `json:"id"`
	LocalityName *string `json:"locality_name"`
	ProvinceName *string `json:"province_name"`
	CountryName  *string `json:"country_name"`
}

type CarriersByLocalityReportResponseDTO struct {
	LocalityID    int    `json:"locality_id"`
	LocalityName  string `json:"locality_name"`
	CarriersCount int    `json:"carriers_count"`
}
