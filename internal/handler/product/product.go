package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/product"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewProductHandler(service service.IProduct) *ProductHandler {
	return &ProductHandler{service: service}
}

type ProductHandler struct {
	service service.IProduct
}

func (h *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		v, err := h.service.GetProduct()
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}
		vResponse := []dto.ProductResponseDTO{}

		for _, prod := range v {

			vResponse = append(vResponse, helper.ProductToProductResponseDTO(prod))

		}
		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vResponse,
		})
	}
}

func (h *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.Error(w, http.StatusBadRequest, errorHandler.INVALID_ID)
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorHandler.INVALID_ID)
			return
		}
		// process
		// - get product by id
		v, err := h.service.GetProductByID(idInt)
		vResponse := helper.ProductToProductResponseDTO(v)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vResponse,
		})
	}
}

func (h *ProductHandler) GetRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id := r.URL.Query().Get("id")
		// id := chi.URLParam(r, "id")
		if id == "" {
			// request
			v, err := h.service.GetProductRecord()
			if err != nil {
				code, msg := errorHandler.HandleError(err)
				response.Error(w, code, msg)
				return
			}
			vResponse := []dto.ProductRecordCountResponseDTO{}
			for _, prod := range v {
				vResponse = append(vResponse, helper.ProductRecordCountToProductRecordCountResponseDTO(prod))
			}
			// response
			response.JSON(w, http.StatusOK, map[string]any{
				"message": "success",
				"data":    vResponse,
			})
			return
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorHandler.INVALID_ID)
			return
		}
		// process
		// - get product by id
		v, err := h.service.GetProductRecordByID(idInt)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.Error(w, code, msg)
			return
		}
		vResponse := []dto.ProductRecordCountResponseDTO{}
		vResponse = append(vResponse, helper.ProductRecordCountToProductRecordCountResponseDTO(v))

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vResponse,
		})
	}
}

func (h *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestDTO dto.ProductRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		request := helper.ProductRequestDTOToProduct(requestDTO)
		// process
		// - create product
		prod, err := h.service.CreateProduct(&request)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product created",
			"data":    helper.ProductToProductResponseDTO(prod),
		})
	}
}

func (h *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.JSON(w, http.StatusBadRequest, errorHandler.INVALID_ID)
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorHandler.INVALID_ID)
			return
		}
		// process
		// - delete product
		err = h.service.DeleteProduct(idInt)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, "")
	}
}

func (h *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, errorHandler.INVALID_ID)
			return
		}

		var requestDTO dto.ProductRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
			response.Error(w, http.StatusBadRequest, errorHandler.INVALID_BODY)
			return
		}

		request := helper.ProductRequestDTOToProduct(requestDTO)

		// process
		// - update product
		v, err := h.service.UpdateProduct(id, &request)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.Error(w, code, msg)
			return
		}
		var vResponse []dto.ProductResponseDTO
		vResponse = append(vResponse, helper.ProductToProductResponseDTO(*v))

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product updated",
			"data":    vResponse,
		})
	}
}
