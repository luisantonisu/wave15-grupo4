package repository

import (
	"database/sql"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type PurchaseOrderRepository struct {
	db *sql.DB
}

func NewPurchaseOrderRepository(defaultDB *sql.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{
		db: defaultDB,
	}
}

// Create a new purchase order
func (r *PurchaseOrderRepository) Create(purchaseOrder model.PurchaseOrderAttributes) (model.PurchaseOrder, error) {
	// Validate order number doesnt already exist
	if r.orderNumberExists(*purchaseOrder.OrderNumber) {
		return model.PurchaseOrder{}, eh.GetErrAlreadyExists(eh.ORDER_NUMBER)
	}
	
	// Create new purchase order in DB
	row, err := r.db.Exec("INSERT INTO purchase_orders (order_number, order_date, tracking_code,  buyer_id, carrier_id, order_status_id, warehouse_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
		purchaseOrder.OrderNumber, purchaseOrder.OrderDate, purchaseOrder.TrackingCode, purchaseOrder.BuyerID, purchaseOrder.CarrierID, purchaseOrder.OrderStatusID, purchaseOrder.WarehouseID,
	)
	
	if err != nil {
		return model.PurchaseOrder{}, eh.GetErrInvalidData(eh.PURCHASE_ORDER)
	}
	id, err := row.LastInsertId()
	if err != nil {
		return model.PurchaseOrder{}, eh.GetErrInvalidData(eh.PURCHASE_ORDER)
	}

	// Response
	var newPurchaseOrder model.PurchaseOrder
	newPurchaseOrder.ID = int(id)
	newPurchaseOrder.PurchaseOrderAttributes = purchaseOrder

	return newPurchaseOrder, nil
}

// Validate if order number is already in use
func (r *PurchaseOrderRepository) orderNumberExists(orderNumber string) bool {
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM purchase_orders WHERE order_number = ?)", orderNumber).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
