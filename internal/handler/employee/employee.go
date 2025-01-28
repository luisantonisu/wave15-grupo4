package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/employee"
)

func NewEmployeeHandler(sv service.IEmployee) *EmployeeHandler {
	return &EmployeeHandler{sv: sv}
}

type EmployeeHandler struct {
	sv service.IEmployee
}

func (h *EmployeeHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e, err := h.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := helper.MapEmployeeToEmployeeDTO(e)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}
