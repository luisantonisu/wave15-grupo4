package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/section"
)

func NewSectionHandler(sv service.ISection) *SectionHandler {
	return &SectionHandler{sv: sv}
}

type SectionHandler struct {
	sv service.ISection
}

func (h *SectionHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sections, err := h.sv.GetAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := make(map[int]dto.SectionResponseDTO)
		for key, value := range sections {
			data[key] = helper.SectionToSectionResponseDTO(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *SectionHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "invalid Id")
			return
		}

		section, err := h.sv.GetByID(id)
		if err != nil {
			if err.Error() == "not exist" {
				response.JSON(w, http.StatusNotFound, "ID not found")
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := helper.SectionToSectionResponseDTO(section)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *SectionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sectionRequest dto.SectionRequestDTO

		err := json.NewDecoder(r.Body).Decode(&sectionRequest)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		section := helper.SectionRequestDTOToSection(sectionRequest)
		section, err = h.sv.Create(section)

		if err != nil {
			if err.Error() == "section number already exists" {
				response.JSON(w, http.StatusConflict, "Section number already exists")
				return
			}

			if err.Error() == "invalid section data" {
				response.JSON(w, http.StatusUnprocessableEntity, "Invalid section data")
				return
			}
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		data := helper.SectionToSectionResponseDTO(section)

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}
