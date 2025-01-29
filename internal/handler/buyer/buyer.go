package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
	"github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewBuyerHandler(sv service.IBuyer) *BuyerHandler {
	return &BuyerHandler{sv: sv}
}

type BuyerHandler struct {
	sv service.IBuyer
}

func (h *BuyerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode body
		var buyerRequestDto dto.BuyerRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&buyerRequestDto); err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// Validate all fields except id
		if err := buyerRequestDto.Validate(); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		newBuyer := helper.BuyerRequestDTOToBuyer(buyerRequestDto)

		// Call service
		data, err := h.sv.Create(newBuyer)
		if err != nil {
			if errors.Is(err, error_handler.CardNumberIdAlreadyInUse) {
				response.Error(w, http.StatusConflict, err.Error())
				return
			}
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Return response
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}

// List all buyers
func (h *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call service
		buyers, err := h.sv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Convert to Response format
		var data []dto.BuyerResponseDTO
		for _, value := range buyers {
			data = append(data, helper.BuyerToBuyerResponseDTO(value))
		}

		// Return response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

// Get a buyer by id
func (h *BuyerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get id from url
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// Call service
		data, err := h.sv.GetByID(id)
		if err != nil {
			if errors.Is(err, error_handler.IDNotFound) {
				response.Error(w, http.StatusNotFound, "Buyer not found")
				return
			}
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Return response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": helper.BuyerToBuyerResponseDTO(data),
		})
	}
}

// Delete a buyer by id
func (h *BuyerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get id from url
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())
			return
		}

		// Call service
		err = h.sv.Delete(id)
		if err != nil {
			if errors.Is(err, error_handler.IDNotFound) {
				response.Error(w, http.StatusNotFound, "Buyer not found")
				return
			}
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Return response
		response.JSON(w, http.StatusNoContent, nil)
	}
}
