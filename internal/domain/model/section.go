package model

type Section struct {
	ID int
	SectionAttributes
}

type SectionAttributes struct {
	SectionNumber      int
	CurrentTemperature float64
	MinimumTemperature float64
	CurrentCapacity    int
	MinimumCapacity    int
	MaximumCapacity    int
	WarehouseID        int
	ProductTypeID      int
}

type SectionAttributesPtr struct {
	SectionNumber      *int
	CurrentTemperature *float64
	MinimumTemperature *float64
	CurrentCapacity    *int
	MinimumCapacity    *int
	MaximumCapacity    *int
	WarehouseID        *int
	ProductTypeID      *int
}

type ReportProductsBatches struct {
	SectionID     int `json:"section_id"`
	SectionNumber int `json:"section_number"`
	ProductsCount int `json:"products_count"`
}
