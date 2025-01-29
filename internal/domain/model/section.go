package model

type Section struct {
	Id int
	SectionAttributes
}

type SectionAttributes struct {
	SectionNumber      string
	CurrentTemperature float64
	MinimumTemperature float64
	CurrentCapacity    int
	MinimumCapacity    int
	MaximumCapacity    int
	WarehouseId        int
	ProductTypeId      int
	ProductBatchId     []int
}
