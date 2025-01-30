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

func (wh *WarehouseHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warehouseRequest dto.WarehouseRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&warehouseRequest); err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		warehouse := helper.WarehouseRequestDTOToWarehouse(warehouseRequest)
		createdWarehouse, err := wh.sv.Create(warehouse)
		if err != nil {
			if err.Error() == "warehouse code is required" {
				response.JSON(w, http.StatusUnprocessableEntity, err.Error())
				return
			}

			if err.Error() == "warehouse code already exists" {
				response.JSON(w, http.StatusConflict, err.Error())
				return
			}

			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		data := helper.WarehouseToWarehouseResponseDTO(createdWarehouse)

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}

func (wh *WarehouseHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid id")
			return
		}

		var warehouseRequest dto.WarehouseRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&warehouseRequest); err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		warehouse := helper.WarehouseRequestDTOToWarehouse(warehouseRequest)
		updatedWarehouse, err := wh.sv.Update(id, warehouse)
		if err != nil {
			if err.Error() == "warehouse not found"{
				response.JSON(w, http.StatusNotFound, err.Error())
				return
			}

			if err.Error() == "warehouse code already exists" {
				response.JSON(w, http.StatusConflict, err.Error())
				return
			}

			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		data := helper.WarehouseToWarehouseResponseDTO(updatedWarehouse)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}
