package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/response"
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

// Ping to check handler status
func (b *BuyerHandler) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

// Create a new buyer
func (b *BuyerHandler) Create() http.HandlerFunc {
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
		data, err := b.sv.Create(newBuyer)
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
func (b *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Call service
		allBuyers, err := b.sv.GetAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		// Convert to Response format
		var data []dto.BuyerResponseDTO
		for _, value := range allBuyers {
			data = append(data, helper.BuyerToBuyerResponseDTO(value))
		}

		// Return response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}
