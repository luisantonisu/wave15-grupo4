package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	buyerRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/buyer"
	carryRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/carry"
	orderStatusRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/order_status"
	purchaseOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/purchase_order"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type PurchaseOrderService struct {
	purchaseOrderRp purchaseOrderRepository.IPurchaseOrder
	buyerRp         buyerRepository.IBuyer
	carryRp         carryRepository.ICarry
	orderStatusRp   orderStatusRepository.IOrderStatus
	warehouseRp     warehouseRepository.IWarehouse
}

func NewPurchaseOrderService(purchaseOrderRp purchaseOrderRepository.IPurchaseOrder, buyerRp buyerRepository.IBuyer, carryRp carryRepository.ICarry, orderStatusRp orderStatusRepository.IOrderStatus, warehouseRp warehouseRepository.IWarehouse) *PurchaseOrderService {
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
