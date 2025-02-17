package dto

import eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"

type SectionResponseDTO struct {
	ID                 int     `json:"id"`
	SectionNumber      int     `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

type SectionRequestDTO struct {
	SectionNumber      int     `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

type SectionRequestDTOPtr struct {
	SectionNumber      *int     `json:"section_number"`
	CurrentTemperature *float64 `json:"current_temperature"`
	MinimumTemperature *float64 `json:"minimum_temperature"`
	CurrentCapacity    *int     `json:"current_capacity"`
	MinimumCapacity    *int     `json:"minimum_capacity"`
	MaximumCapacity    *int     `json:"maximum_capacity"`
	WarehouseID        *int     `json:"warehouse_id"`
	ProductTypeID      *int     `json:"product_type_id"`
}

type ReportProductsBatchesResponseDTO struct {
	SectionID     int `json:"section_id"`
	SectionNumber int `json:"section_number"`
	ProductsCount int `json:"products_count"`
}

func (e *SectionRequestDTO) Validate() error {
	if e.SectionNumber <= 0 {
		return eh.GetErrInvalidData(eh.SECTION)
	}
	return nil
}
