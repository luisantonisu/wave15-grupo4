package handler

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
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
