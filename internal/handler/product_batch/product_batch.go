package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/product_batch"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewProductBatchHandler(pb service.IProductBatch) *ProductBatchHandler {
	return &ProductBatchHandler{pb: pb}
}

type ProductBatchHandler struct {
	pb service.IProductBatch
}

func (h *ProductBatchHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var productBatchRequest dto.ProductBatchRequestDTO

		err := json.NewDecoder(r.Body).Decode(&productBatchRequest)
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}

		productBatch := helper.ProductBatchRequestDTOToProductBatch(productBatchRequest)
		productBatch, err = h.pb.Create(productBatch)

		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := helper.ProductBatchToProductBatchResponseDTO(productBatch)

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}
