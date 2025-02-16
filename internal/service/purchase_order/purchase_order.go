package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	buyer_rp "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	carry_rp "github.com/luisantonisu/wave15-grupo4/internal/repository/carry"
	order_status_rp "github.com/luisantonisu/wave15-grupo4/internal/repository/order_status"
	purchase_order_rp "github.com/luisantonisu/wave15-grupo4/internal/repository/purchase_order"
	warehouse_rp "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type PurchaseOrderService struct {
	purchaseOrderRp purchase_order_rp.IPurchaseOrder
	buyerRp         buyer_rp.IBuyer
	carryRp         carry_rp.ICarry
	orderStatusRp   order_status_rp.IOrderStatus
	warehouseRp     warehouse_rp.IWarehouse
}

func NewPurchaseOrderService(purchaseOrderRp purchase_order_rp.IPurchaseOrder, buyerRp buyer_rp.IBuyer, carryRp carry_rp.ICarry, orderStatusRp order_status_rp.IOrderStatus, warehouseRp warehouse_rp.IWarehouse) *PurchaseOrderService {
	return &PurchaseOrderService{
		purchaseOrderRp: purchaseOrderRp,
		buyerRp:         buyerRp,
		carryRp:         carryRp,
		orderStatusRp:   orderStatusRp,
		warehouseRp:     warehouseRp,
	}
}

// Create new purchase order
func (s *PurchaseOrderService) Create(purchaseOrder model.PurchaseOrderAttributes) (model.PurchaseOrder, error) {
	// Validate Buyer exist
	_, err := s.buyerRp.GetByID(purchaseOrder.BuyerID)
	if err != nil {
		return model.PurchaseOrder{}, eh.GetErrForeignKey(eh.BUYER)
	}

	// Validate Order Status exist
	_, err = s.orderStatusRp.GetByID(purchaseOrder.OrderStatusID)
	if err != nil {
		return model.PurchaseOrder{}, eh.GetErrForeignKey(eh.ORDER_STATUS)
	}

	// Validate Warehouse exist
	_, err = s.warehouseRp.GetByID(purchaseOrder.WarehouseID)
	if err != nil {
		return model.PurchaseOrder{}, eh.GetErrForeignKey(eh.WAREHOUSE)
	}

	// Validate Carrier exist
	_, err = s.carryRp.GetByID(purchaseOrder.CarrierID)
	if err != nil {
		return model.PurchaseOrder{}, eh.GetErrForeignKey(eh.CARRY)
	}

	return s.purchaseOrderRp.Create(purchaseOrder)
}
