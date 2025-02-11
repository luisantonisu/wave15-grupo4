package model

type InboundOrder struct {
	ID int
	InboundOrderAttributes
}

type InboundOrderAttributes struct {
	OrderDate      string
	OrderNumber    string
	EmployeeID     int
	ProductBatchID int
	WarehouseID    int
}

type InboundOrderAttributesPtr struct {
	OrderDate      *string
	OrderNumber    *string
	EmployeeID     *int
	ProductBatchID *int
	WarehouseID    *int
}
