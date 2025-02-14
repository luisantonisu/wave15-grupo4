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
	if inboundOrder.OrderDate == "" || inboundOrder.EmployeeID <= 0 || inboundOrder.OrderNumber <= 0 || inboundOrder.WarehouseID <= 0 || inboundOrder.ProductBatchID <= 0 {
		return model.InboundOrder{}, eh.GetErrInvalidData(eh.INBOUND_ORDER)
	}

	_, err := h.employeeRp.GetByID(inboundOrder.EmployeeID)
	if err != nil {
		return model.InboundOrder{}, eh.GetErrForeignKey(eh.EMPLOYEE)
	}

	_, err = h.warehouseRp.GetByID(inboundOrder.WarehouseID)
	if err != nil {
		return model.InboundOrder{}, eh.GetErrForeignKey(eh.WAREHOUSE)
	}

	return h.ibOrdRp.CreateInboundOrder(inboundOrder)
}
