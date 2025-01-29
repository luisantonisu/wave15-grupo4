package handler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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

func (h *SellerHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, nil)
			return
		}

		//process
		seller, err := h.sv.GetByID(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"status_code": http.StatusNotFound, 
				"message" : err.Error(),
			})
			return
		}

		//mapping
		var data = helper.SellerToSellerResponseDTO(seller)

		//response
		var resp []dto.SellerResponseDTO
		response.JSON(w, http.StatusOK, map[string]any{
			"data": append(resp, data),
		})
	}
}
