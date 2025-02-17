package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/buyer"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}

		// Validate card number id is present
		if err := buyerRequestDto.Validate(); err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}
		newBuyerRequest := helper.BuyerRequestDTOToBuyerAttributes(buyerRequestDto)

		// Call service
		newBuyer, err := h.sv.Create(newBuyerRequest)
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}
		data := helper.BuyerToBuyerResponseDTO(newBuyer)

		// Return response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// List all buyers
func (h *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call service
		buyers, err := h.sv.GetAll()
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
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
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		// Call service
		buyer, err := h.sv.GetByID(id)
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}
		data := helper.BuyerToBuyerResponseDTO(buyer)

		// Return response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
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
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		// Call service
		err = h.sv.Delete(id)
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}

		// Return response
		response.JSON(w, http.StatusNoContent, nil)
	}
}

// Update a buyer info by id
func (h *BuyerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get id from url
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		// Decode body
		var buyerRequestDto dto.BuyerRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&buyerRequestDto); err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}
		newBuyer := helper.BuyerRequestDTOToBuyer(buyerRequestDto)

		// Call service
		updatedBuyer, err := h.sv.Update(id, newBuyer)
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}
		data := helper.BuyerToBuyerResponseDTO(updatedBuyer)

		// Return response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// Generate Purchase Order Report
func (h *BuyerHandler) Report() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode query params
		var id int
		var err error
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			id = -1
		} else {
			id, err = strconv.Atoi(idStr)
			if err != nil {
				response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
				return
			}
		}

		// Call service
		report, err := h.sv.PurchaseOrderReport(id)
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}

		// Return response
		data := []dto.ReportPurchaseOrdersResponseDTO{}
		for _, buyer := range report {
			data = append(data, helper.ReportPurchaseOrdersToReportPurchaseOrdersResponseDTO(buyer))
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
