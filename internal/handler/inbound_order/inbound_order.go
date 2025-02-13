package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/inbound_order"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type InboundOrderHandler struct {
	sv service.IInboundOrder
}

func NewInboundOrderHandler(sv service.IInboundOrder) *InboundOrderHandler {
	return &InboundOrderHandler{sv: sv}
}

func (h *InboundOrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inbDto dto.InboundOrderRequestDTO

		if err := json.NewDecoder(r.Body).Decode(&inbDto); err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		inboundOrder := helper.InboundOrderRequestDTOToInboundOrder(inbDto)
		inbOrd, err := h.sv.Create(inboundOrder)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := helper.InboundOrderToInboundOrderResponseDTO(inbOrd)

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}
