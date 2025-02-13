package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/purchase_order"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewPurchaseOrderHandler(sv service.IPurchaseOrder) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{sv: sv}
}

type PurchaseOrderHandler struct {
	sv service.IPurchaseOrder
}

func (h *PurchaseOrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode body
		var purchaseOrderRequestDto dto.PurchaseOrderRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&purchaseOrderRequestDto); err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}

		// Validate order number id is present
		if err := purchaseOrderRequestDto.Validate(); err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}
		newPurchaseOrder := helper.PurchaseOrderRequestDTOToPurchaseOrderAttributes(purchaseOrderRequestDto)

		// Call service
		data, err := h.sv.Create(newPurchaseOrder)
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}

		// Return response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Purchase Order created",
			"data":    data,
		})
	}
}
