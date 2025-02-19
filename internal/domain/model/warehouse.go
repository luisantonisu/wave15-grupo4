package model

type Warehouse struct {
	ID int
	WarehouseAttributes
}

type WarehouseAttributes struct {
	WarehouseCode      *string
	Address            *string
	Telephone          *uint
	MinimumCapacity    *int
	MinimumTemperature *float32
	LocalityID         *int
}
