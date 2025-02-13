package dto

type InboundOrderRequestDTO struct {
	OrderDate      string `json:"order_date"`
	OrderNumber    int    `json:"order_number"`
	EmployeeID     int    `json:"employee_id"`
	ProductBatchID int    `json:"product_batch_id"`
	WarehouseID    int    `json:"warehouse_id"`
}

type InboundOrderResponseDTO struct {
	ID             int    `json:"id"`
	OrderDate      string `json:"order_date"`
	OrderNumber    int    `json:"order_number"`
	EmployeeID     int    `json:"employee_id"`
	ProductBatchID int    `json:"product_batch_id"`
	WarehouseID    int    `json:"warehouse_id"`
}

type InboundOrderRequestDTOPtr struct {
	OrderDate      *string `json:"order_date"`
	OrderNumber    *int    `json:"order_number"`
	EmployeeID     *int    `json:"employee_id"`
	ProductBatchID *int    `json:"product_batch_id"`
	WarehouseID    *int    `json:"warehouse_id"`
}
