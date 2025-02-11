package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func InboundOrderToInboundOrderResponseDTO(inboundOrder map[int]model.InboundOrder) map[int]dto.InboundOrderResponseDTO {
	data := make(map[int]dto.InboundOrderResponseDTO)
	for _, value := range inboundOrder {
		data[value.ID] = dto.InboundOrderResponseDTO{
			ID:             value.ID,
			OrderDate:      value.InboundOrderAttributes.OrderDate,
			OrderNumber:    value.InboundOrderAttributes.OrderNumber,
			EmployeeID:     value.InboundOrderAttributes.EmployeeID,
			ProductBatchID: value.InboundOrderAttributes.ProductBatchID,
			WarehouseID:    value.InboundOrderAttributes.WarehouseID,
		}
	}
	return data
}

func InboundOrderRequestDTOToInboundOrder(inboundOrder dto.InboundOrderRequestDTO) model.InboundOrderAttributes {
	data := model.InboundOrderAttributes{
		OrderDate:      inboundOrder.OrderDate,
		OrderNumber:    inboundOrder.OrderNumber,
		EmployeeID:     inboundOrder.EmployeeID,
		ProductBatchID: inboundOrder.ProductBatchID,
		WarehouseID:    inboundOrder.WarehouseID,
	}
	return data
}

func InboundOrderRequestDTOPtrToInboundOrderPtr(inboundOrder dto.InboundOrderRequestDTOPtr) model.InboundOrderAttributesPtr {
	data := model.InboundOrderAttributesPtr{
		OrderDate:      inboundOrder.OrderDate,
		OrderNumber:    inboundOrder.OrderNumber,
		EmployeeID:     inboundOrder.EmployeeID,
		ProductBatchID: inboundOrder.ProductBatchID,
		WarehouseID:    inboundOrder.WarehouseID,
	}
	return data
}
