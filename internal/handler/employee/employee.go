package handler

import (
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
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

func (h *EmployeeHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid id")
			return
		}

		e, err := h.sv.GetByID(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		data := helper.EmployeeToEmployeeDTO(e)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}
