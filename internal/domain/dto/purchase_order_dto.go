package dto

import eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"

type PurchaseOrderRequestDTO struct {
	OrderNumber   string `json:"order_number"`
	OrderDate     string `json:"order_date"`
	TrackingCode  string `json:"tracking_code"`
	BuyerID       int    `json:"buyer_id"`
	CarrierID     int    `json:"carrier_id"`
	OrderStatusID int    `json:"order_status_id"`
	WarehouseID   int    `json:"warehouse_id"`
}

type PurchaseOrderResponseDTO struct {
	ID            int    `json:"id"`
	PurchaseOrderRequestDTO
}

// Validate PurchaseOrderRequestDTO required fields
func (e *PurchaseOrderRequestDTO) Validate() error {
	if e.OrderNumber == "" {
		return eh.GetErrInvalidData(eh.ORDER_NUMBER)
	}
	return nil
}
