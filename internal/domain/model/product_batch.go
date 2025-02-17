package model

type ProductBatch struct {
	ID int
	ProductBatchAttributes
}

type ProductBatchAttributes struct {
	BatchNumber        string
	CurrentQuantity    int
	CurrentTemperature float64
	DueDate            string
	InitialQuantity    int
	ManufacturingDate  string
	ManufacturingHour  string
	MinimumTemperature float64
	ProductID          int
	SectionID          int
}
