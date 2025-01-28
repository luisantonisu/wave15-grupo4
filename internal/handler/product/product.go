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

func NewProductHandler(sv service.IProduct) *ProductHandler {
	return &ProductHandler{sv: sv}
}

type ProductHandler struct {
	sv service.IProduct
}

func (p *ProductHandler) GetProductsHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := p.sv.GetProduct()
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

func (p *ProductHandler) GetProductByIdHTTP() http.HandlerFunc {
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
		// - get all vehicles
		v, err := p.sv.GetProductById(idInt)
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
func (p *ProductHandler) CreateProductHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...
		var requestDTO dto.ProductDTO
		if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		request := helper.MapProductDTOToProduct(requestDTO)
		// process
		// - get all vehicles
		v, err := p.sv.CreateProduct(&request)
		if err != nil {
			response.JSON(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}
