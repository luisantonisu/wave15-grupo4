package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/product_record"
	errorHandler "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

type ProductRecordHandler struct {
	service service.IProductRecord
}

func NewProductRecordHandler(service service.IProductRecord) *ProductRecordHandler {
	return &ProductRecordHandler{service: service}
}

func (productRecordHandler *ProductRecordHandler) GetByID() http.HandlerFunc {
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

		v, err := productRecordHandler.service.GetProductRecordByID(idInt)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}

func (productRecordHandler *ProductRecordHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestDTO dto.ProductRecordRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		request := helper.ProductRecordRequestDTOToProductRecord(requestDTO)
		// process
		// - create product
		err := productRecordHandler.service.CreateProductRecord(request)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product record created",
		})
	}
}
