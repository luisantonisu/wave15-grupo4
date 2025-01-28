package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
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
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})
	}
}
