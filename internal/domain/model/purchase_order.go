package model

type PurchaseOrder struct {
	ID int
	PurchaseOrderAttributes
}

type PurchaseOrderAttributes struct {
	OrderNumber   *string
	OrderDate     *string
	TrackingCode  *string
	BuyerID       *int
	CarrierID     *int
	OrderStatusID *int
	WarehouseID   *int
}
