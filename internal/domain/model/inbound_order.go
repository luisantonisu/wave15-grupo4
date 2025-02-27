package model

type InboundOrder struct {
	ID int
	InboundOrderAttributes
}

type InboundOrderAttributes struct {
	OrderDate      *string
	OrderNumber    *int
	EmployeeID     *int
	ProductBatchID *int
	WarehouseID    *int
}
