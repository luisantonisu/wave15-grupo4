package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	employeeRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/employee"
	inboundOrderRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/inbound_order"
	warehouseRepository "github.com/luisantonisu/wave15-grupo4/internal/repository/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type InboundOrderService struct {
	ibOrdRp     inboundOrderRepository.IInboundOrder
	employeeRp  employeeRepository.IEmployee
	warehouseRp warehouseRepository.IWarehouse
}

func NewInboundOrderService(ibOrdRp inboundOrderRepository.IInboundOrder, employeeRp employeeRepository.IEmployee, warehouseRp warehouseRepository.IWarehouse) *InboundOrderService {
	return &InboundOrderService{
		ibOrdRp:     ibOrdRp,
		employeeRp:  employeeRp,
		warehouseRp: warehouseRp,
	}
}

func (h *InboundOrderService) Create(inboundOrder model.InboundOrderAttributes) (model.InboundOrder, error) {
	if inboundOrder.OrderDate == nil || inboundOrder.EmployeeID == nil || inboundOrder.OrderNumber == nil || inboundOrder.WarehouseID == nil || inboundOrder.ProductBatchID == nil || *inboundOrder.OrderDate == ""{
		return model.InboundOrder{}, eh.GetErrInvalidData(eh.INBOUND_ORDER)
	}

	_, err := h.employeeRp.GetByID(*inboundOrder.EmployeeID)
	if err != nil {
		return model.InboundOrder{}, eh.GetErrForeignKey(eh.EMPLOYEE)
	}

	_, err = h.warehouseRp.GetByID(*inboundOrder.WarehouseID)
	if err != nil {
		return model.InboundOrder{}, eh.GetErrForeignKey(eh.WAREHOUSE)
	}

	if !h.ibOrdRp.AlreadyExists("product_batch_id", *inboundOrder.ProductBatchID) {
		return model.InboundOrder{}, eh.GetErrForeignKey(eh.PRODUCT_BATCH_ID)
	}

	if h.ibOrdRp.AlreadyExists("order_number", *inboundOrder.OrderNumber) {
		return model.InboundOrder{}, eh.GetErrAlreadyExistsCompose(eh.INBOUND_ORDER, eh.ORDER_NUMBER)
	}

	return h.ibOrdRp.CreateInboundOrder(inboundOrder)
}
