package handler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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

		var data []dto.WarehouseResponseDTO
		for _, value := range warehouses {
			data = append(data, helper.WarehouseToWarehouseResponseDTO(value))
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (wh *WarehouseHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid id")
			return
		}

		warehouse, err := wh.sv.GetByID(id)
		if err != nil {
			if err.Error() == "warehouse not found" ||
				err.Error() == "no warehouses found" {
				response.JSON(w, http.StatusNotFound, err.Error())
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := helper.WarehouseToWarehouseResponseDTO(warehouse)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}
