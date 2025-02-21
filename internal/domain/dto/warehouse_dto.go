package dto

type WarehouseResponseDTO struct {
	ID int `json:"id"`
	WarehouseRequestDTO
}

type WarehouseRequestDTO struct {
	WarehouseCode      *string  `json:"warehouse_code"`
	Address            *string  `json:"address"`
	Telephone          *uint    `json:"telephone"`
	MinimumCapacity    *int     `json:"minimum_capacity"`
	MinimumTemperature *float32 `json:"minimum_temperature"`
	LocalityID         *int     `json:"locality_id"`
}
