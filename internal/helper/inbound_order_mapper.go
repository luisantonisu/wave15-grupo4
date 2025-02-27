package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func InboundOrderToInboundOrderResponseDTO(inboundOrder model.InboundOrder) dto.InboundOrderResponseDTO {
	return dto.InboundOrderResponseDTO{
		ID:             inboundOrder.ID,
		InboundOrderRequestDTO: dto.InboundOrderRequestDTO{
			OrderDate:      inboundOrder.InboundOrderAttributes.OrderDate,
			OrderNumber:    inboundOrder.InboundOrderAttributes.OrderNumber,
			EmployeeID:     inboundOrder.InboundOrderAttributes.EmployeeID,
			ProductBatchID: inboundOrder.InboundOrderAttributes.ProductBatchID,
			WarehouseID:    inboundOrder.InboundOrderAttributes.WarehouseID,
		},
	}
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
