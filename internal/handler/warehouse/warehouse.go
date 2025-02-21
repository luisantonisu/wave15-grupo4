package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/warehouse"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewWarehouseHandler(sv service.IWarehouse) *WarehouseHandler {
	return &WarehouseHandler{sv: sv}
}

type WarehouseHandler struct {
	sv service.IWarehouse
}

func (h *WarehouseHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		warehouses, err := h.sv.GetAll()
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := []dto.WarehouseResponseDTO{}
		for _, value := range warehouses {
			data = append(data, helper.WarehouseToWarehouseResponseDTO(value))
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *WarehouseHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		warehouse, err := h.sv.GetByID(id)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := helper.WarehouseToWarehouseResponseDTO(warehouse)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *WarehouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouseRequest dto.WarehouseRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&warehouseRequest); err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}

		warehouse := helper.WarehouseRequestDTOToWarehouse(warehouseRequest)
		createdWarehouse, err := h.sv.Create(warehouse)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := helper.WarehouseToWarehouseResponseDTO(createdWarehouse)

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "warehouse created",
			"data":    data,
		})
	}
}

func (h *WarehouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		var warehouseRequest dto.WarehouseRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&warehouseRequest); err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}

		warehouse := helper.WarehouseRequestDTOToWarehouse(warehouseRequest)
		updatedWarehouse, err := h.sv.Update(id, warehouse.WarehouseAttributes)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := helper.WarehouseToWarehouseResponseDTO(updatedWarehouse)

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "warehouse updated",
			"data":    data,
		})
	}
}

func (h *WarehouseHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}
