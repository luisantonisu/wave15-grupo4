package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"
)

func NewWarehouseHandler(sv service.IWarehouse) *WarehouseHandler {
	return &WarehouseHandler{sv: sv}
}

type WarehouseHandler struct {
	sv service.IWarehouse
}

func (wh *WarehouseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := wh.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := make(map[int]dto.WarehouseResponseDTO)
		for key, value := range warehouses {
			data[key] = helper.WarehouseToWarehouseResponseDTO(value)
		}
		response.JSON(w, http.StatusOK, data)
	}
}
