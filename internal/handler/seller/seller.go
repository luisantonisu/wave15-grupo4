package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/seller"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
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
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		//process
		seller, err := h.sv.GetByID(id)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		//mapping
		var data = helper.SellerToSellerResponseDTO(seller)

		//response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *SellerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		var newSeller dto.SellerRequestDTO
		json.NewDecoder(r.Body).Decode(&newSeller)

		//mapping
		var seller = helper.SellerRequestDTOToSeller(newSeller)

		//process
		result, err := h.sv.Create(seller)
		if err != nil {

			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return

		}

		//mapping
		var data = helper.SellerToSellerResponseDTO(result)

		//response
		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}

func (h *SellerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		var updateSeller dto.SellerRequestDTO
		json.NewDecoder(r.Body).Decode(&updateSeller)

		//mapping
		var seller = helper.SellerRequestDTOPtrToSellerPtr(updateSeller)

		//process
		result, err := h.sv.Update(id, seller)
		if err != nil {

			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return

		}

		//mapping
		var data = helper.SellerToSellerResponseDTO(result)

		//response
		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *SellerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		//process
		err = h.sv.Delete(id)
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}

		//Response
		response.JSON(w, http.StatusNoContent, nil)
	}
}
