package model

type Warehouse struct {
	Id                 int     `json:"id"`
	WarehouseAttributes
}

type WarehouseAttributes struct {
	WarehouseCode      string     `json:"warehouse_code"`
	Address            string  `json:"address"`
	Telephone          uint    `json:"telephone"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float32 `json:"minimum_temperature"`
}