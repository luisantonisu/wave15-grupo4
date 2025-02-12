package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/inbound_order"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type InboundOrderService struct {
	rp repository.IInboundOrder
}

func NewInboundOrderService(rp repository.IInboundOrder) *InboundOrderService {
	return &InboundOrderService{rp: rp}
}

func (h *InboundOrderService) Create(inboundOrder model.InboundOrderAttributes) (model.InboundOrder, error) {
	if inboundOrder.OrderDate == "" || inboundOrder.EmployeeID <= 0 || inboundOrder.OrderNumber <= 0 || inboundOrder.WarehouseID <= 0 || inboundOrder.ProductBatchID <= 0 {
		return model.InboundOrder{}, eh.GetErrInvalidData(eh.INBOUND_ORDER)
	}
	return h.rp.CreateInboundOrder(inboundOrder)
}
