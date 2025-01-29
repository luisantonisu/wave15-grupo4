package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/seller"
)

func NewSellerHandler(sv service.ISeller) *SellerHandler {
	return &SellerHandler{sv: sv}
}

type SellerHandler struct {
	sv service.ISeller
}

func (h *SellerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//process
		sellers, err := h.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		//mapping
		var data []dto.SellerResponseDTO
		for _, value := range sellers {
			data = append(data, helper.SellerToSellerResponseDTO(value))
		}

		//response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})

	}
}
