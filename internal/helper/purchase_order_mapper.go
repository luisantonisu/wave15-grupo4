package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func PurchaseOrderRequestDTOToPurchaseOrderAttributes(purchaseOrderRequestDto dto.PurchaseOrderRequestDTO) model.PurchaseOrderAttributes {
	data := model.PurchaseOrderAttributes{
		OrderNumber:   purchaseOrderRequestDto.OrderNumber,
		OrderDate:     purchaseOrderRequestDto.OrderDate,
		TrackingCode:  purchaseOrderRequestDto.TrackingCode,
		BuyerID:       purchaseOrderRequestDto.BuyerID,
		CarrierID:     purchaseOrderRequestDto.CarrierID,
		OrderStatusID: purchaseOrderRequestDto.OrderStatusID,
		WarehouseID:   purchaseOrderRequestDto.WarehouseID,
	}
	return data
}

func PurchaseOrderToPurchaseOrderResponseDTO(purchaseOrder model.PurchaseOrder) dto.PurchaseOrderResponseDTO {
	data := dto.PurchaseOrderResponseDTO{
		ID:            purchaseOrder.ID,
		PurchaseOrderRequestDTO: dto.PurchaseOrderRequestDTO{
			OrderNumber:   purchaseOrder.OrderNumber,
			OrderDate:     purchaseOrder.OrderDate,
			TrackingCode:  purchaseOrder.TrackingCode,
			BuyerID:       purchaseOrder.BuyerID,
			CarrierID:     purchaseOrder.CarrierID,
			OrderStatusID: purchaseOrder.OrderStatusID,
			WarehouseID:   purchaseOrder.WarehouseID,
		},
	}
	return data
}
