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
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
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
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		var data []dto.SectionResponseDTO
		for _, value := range sections {
			data = append(data, helper.SectionToSectionResponseDTO(value))
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
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		section, err := h.sv.GetByID(id)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
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
		var sectionRequestDto dto.SectionRequestDTO

		err := json.NewDecoder(r.Body).Decode(&sectionRequestDto)
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}

		err = sectionRequestDto.Validate()
		if err != nil {
			code, message := eh.HandleError(err)
			response.Error(w, code, message)
			return
		}
		newSectionRequest := helper.SectionRequestDTOToSection(sectionRequestDto)

		newSection, err := h.sv.Create(newSectionRequest)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := helper.SectionToSectionResponseDTO(newSection)

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": data,
		})
	}
}

func (h *SectionHandler) Patch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		var sectionRequestDTO dto.SectionRequestDTO
		err = json.NewDecoder(r.Body).Decode(&sectionRequestDTO)
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_BODY)
			return
		}
		section := helper.SectionRequestDTOPtrToSectionPtr(sectionRequestDTO)

		updatedSection, err := h.sv.Patch(id, section)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}
		data := helper.SectionToSectionResponseDTO(updatedSection)

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}

func (h *SectionHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			return
		}

		err = h.sv.Delete(id)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

func (h *SectionHandler) Report() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id *int
		var err error
		idStr := r.URL.Query()
		if idStr.Has("id") {
			hasId, err := strconv.Atoi(idStr.Get("id"))
			if err != nil {
				response.Error(w, http.StatusBadRequest, eh.INVALID_ID)
			}
			id = &hasId
		}

		report, err := h.sv.Report(id)
		if err != nil {
			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}

		data := []dto.ReportProductsBatchesResponseDTO{}
		for _, value := range report {
			data = append(data, helper.ReportProductsBatchesToReportProductsBatchesResponseDTO(value))
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": data,
		})
	}
}
