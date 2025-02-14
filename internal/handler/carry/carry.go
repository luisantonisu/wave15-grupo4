package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/carry"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewCarryHandler(sv service.ICarry) *CarryHandler {
	return &CarryHandler{sv: sv}
}

type CarryHandler struct {
	sv service.ICarry
}

func (h *CarryHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var carryRequest dto.CarryRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&carryRequest); err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}

		carry := helper.CarryRequestDTOToCarry(carryRequest)
		createdCarry, err := h.sv.Create(carry)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := helper.CarryToCarryResponseDTO(createdCarry)

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "carry created",
			"data":    data,
		})
	}
}
