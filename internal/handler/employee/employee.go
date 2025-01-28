package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
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

func (h *EmployeeHandler) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var empDto dto.EmpSaveDTO

		if err := json.NewDecoder(r.Body).Decode(&empDto); err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		var employee model.Employee

		employee.CardNumberId = empDto.CardNumberId
		employee.FirstName = empDto.FirstName
		employee.LastName = empDto.LastName
		employee.WarehouseId = empDto.WarehouseId

		emp, err := h.sv.Save(employee)

		if err != nil && err.Error() == "Card number ID already exists" {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		if err != nil && err.Error() == "Invalid employee data" {
			response.JSON(w, http.StatusUnprocessableEntity, err.Error())
			return
		}

		data := helper.EmployeeToEmployeeDTO(emp)

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}
