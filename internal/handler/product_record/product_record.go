package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
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

func (h *ProductRecordHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var requestDTO dto.ProductRecordRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
			response.Error(w, http.StatusBadRequest, errorHandler.INVALID_BODY)
			return
		}

		request := helper.ProductRecordRequestDTOToProductRecord(requestDTO)
		// process
		// - create product record
		pr, err := h.service.CreateProductRecord(request)
		if err != nil {
			code, msg := errorHandler.HandleError(err)
			response.JSON(w, code, msg)
			return
		}

		// response
		response.JSON(w, http.StatusCreated, map[string]any{
			"data":    helper.ProductRecordToProductRecordResponseDTO(pr),
			"message": "Product record created",
		})
	}
}
