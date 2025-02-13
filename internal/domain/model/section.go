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
	// ProductBatchID     []int
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
	// ProductBatchID     *[]int
}
