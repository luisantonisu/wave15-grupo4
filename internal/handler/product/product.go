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
		// ...

		// process
		// - get all products
		v, err := productHandler.service.GetProduct()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

func (productHandler *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.JSON(w, http.StatusBadRequest, "ID cant be empty or 0")
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
		}
		// process
		// - get product by id
		v, err := productHandler.service.GetProductByID(idInt)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

func (productHandler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		var requestDTO dto.ProductRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		request := helper.ProductRequestDTOToProduct(requestDTO)
		// process
		// - create product
		err := productHandler.service.CreateProduct(&request)
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, err.Error())
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
		// ...
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.JSON(w, http.StatusBadRequest, "ID cant be empty or 0")
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
		}
		// process
		// - delete product
		err = productHandler.service.DeleteProduct(idInt)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		// response
		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "Product deleted",
		})
	}
}

func (productHandler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		id := chi.URLParam(r, "id")
		if id == "" || id == "0" {
			response.JSON(w, http.StatusBadRequest, "ID cant be empty or 0")
			return
		}
		var requestDTO dto.ProductRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		request := helper.ProductRequestDTOToProductPtr(requestDTO)

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
		}
		// process
		// - update product
		product, err := productHandler.service.UpdateProduct(idInt, &request)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product updated",
			"data":    product,
		})
	}
}
