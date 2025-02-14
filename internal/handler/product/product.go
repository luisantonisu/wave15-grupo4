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

func (productHandler *ProductHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		v, err := productHandler.service.GetProduct()
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}
		vResponse := make(map[int]dto.ProductResponseDTO)
		for i, prod := range v {
			vResponse[i] = helper.ProductToProductResponseDTO(prod)
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vResponse,
		})
	}
}

func (productHandler *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.JSON(w, http.StatusBadRequest, errorHandler.INVALID_ID)
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}
		// process
		// - get product by id
		v, err := productHandler.service.GetProductByID(idInt)
		vResponse := helper.ProductToProductResponseDTO(v)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vResponse,
		})
	}
}

func (productHandler *ProductHandler) GetRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id := r.URL.Query().Get("id")
		if id == "" {
			// process with id not found
			v, err := productHandler.service.GetProductRecord()
			if err != nil {
				code, msg := errorHandler.HandleError(err)
				response.JSON(w, code, msg)
				return
			}
			vResponse := make(map[int]dto.ProductRecordCountResponseDTO)
			for i, prod := range v {
				vResponse[i] = helper.ProductRecordCountToProductRecordCountResponseDTO(prod)
			}
			// response
			response.JSON(w, http.StatusOK, map[string]any{
				"message": "success",
				"data":    vResponse,
			})
			return
		}
		// process with id found
		idInt, err := strconv.Atoi(id)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}
		// process
		// - get product by id
		v, err := productHandler.service.GetProductRecordByID(idInt)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}
		vResponse := helper.ProductRecordCountToProductRecordCountResponseDTO(v)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    vResponse,
		})
	}
}

func (productHandler *ProductHandler) Create() http.HandlerFunc {
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
		err := productHandler.service.CreateProduct(&request)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product created",
		})
	}
}

func (productHandler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.JSON(w, http.StatusBadRequest, errorHandler.INVALID_ID)
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}
		// process
		// - delete product
		err = productHandler.service.DeleteProduct(idInt)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, "")
	}
}

func (productHandler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		var requestDTO dto.ProductRequestDTOPtr
		if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
			response.Error(w, http.StatusBadRequest, errorHandler.INVALID_BODY)
			return
		}

		request := helper.ProductRequestDTOPtrToProductPtr(requestDTO)

		// process
		// - update product
		v, err := productHandler.service.UpdateProduct(id, &request)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}
		vResponse := helper.ProductToProductResponseDTO(*v)

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product updated",
			"data":    vResponse,
		})
	}
}
