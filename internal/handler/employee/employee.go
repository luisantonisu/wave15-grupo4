package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
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
		employees, err := h.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := []dto.EmployeeResponseDTO{}
		for _, employee := range employees {
			data = append(data, helper.EmployeeToEmployeeResponseDTO(employee))
		}

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

		employee, err := h.sv.GetByID(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		data := helper.EmployeeToEmployeeResponseDTO(employee)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *EmployeeHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var empDto dto.EmployeeRequestDTO

		if err := json.NewDecoder(r.Body).Decode(&empDto); err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		employee := helper.EmployeeRequestDTOToEmployee(empDto)

		emp, err := h.sv.Create(employee)

		if err != nil {
			if err.Error() == "Card number ID already exists" {
				response.JSON(w, http.StatusBadRequest, err.Error())
			} else if err.Error() == "Invalid employee data" {
				response.JSON(w, http.StatusUnprocessableEntity, err.Error())
			}
			return
		}

		data := helper.EmployeeToEmployeeResponseDTO(emp)

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}

func (h *EmployeeHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid id")
			return
		}

		var empDto dto.EmployeeRequestDTO
		if err := json.NewDecoder(r.Body).Decode(&empDto); err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		employee := helper.EmployeeRequestDTOToEmployee(empDto)

		empUpdated, err := h.sv.Update(id, employee)
		if err != nil {
			if err.Error() == "Employee not found" {
				response.JSON(w, http.StatusNotFound, err.Error())
			} else if err.Error() == "Card number ID already exists" {
				response.JSON(w, http.StatusBadRequest, err.Error())
			}
			return
		}

		data := helper.EmployeeToEmployeeResponseDTO(empUpdated)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *EmployeeHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid id")
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			response.JSON(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}
