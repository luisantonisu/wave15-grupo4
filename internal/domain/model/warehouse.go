package model

import "errors"

var (
	ErrWarehouseCodeEmpty           = errors.New("warehouse code is empty")
	ErrAddressEmpty                 = errors.New("address is empty")
	ErrTelephoneEmptyOrInvalid      = errors.New("telephone is empty or invalid")
	ErrMinimumCapacityNegative      = errors.New("minimum capacity is negative")
	ErrMinimumTemperatureOutOfRange = errors.New("minimum temperature is out of range")
)

type Warehouse struct {
	ID int
	WarehouseAttributes
}

type WarehouseAttributes struct {
	WarehouseCode      string
	Address            string
	Telephone          uint
	MinimumCapacity    int
	MinimumTemperature float32
	LocalityID         int
}

type WarehouseAttributesPtr struct {
	WarehouseCode      *string
	Address            *string
	Telephone          *uint
	MinimumCapacity    *int
	MinimumTemperature *float32
	LocalityID         int
}
