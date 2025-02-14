package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bootcamp-go/web/response"

	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/helper"
	service "github.com/luisantonisu/wave15-grupo4/internal/service/locality"
	eh "github.com/luisantonisu/wave15-grupo4/pkg/error_handler"
)

func NewLocalityHandler(sv service.ILocality) *LocalityHandler {
	return &LocalityHandler{sv: sv}
}

type LocalityHandler struct {
	sv service.ILocality
}

func (h *LocalityHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		var newLocality dto.LocalityRequestDTO
		json.NewDecoder(r.Body).Decode(&newLocality)

		//mapping
		var locality = helper.LocalityRequestDTOToLocality(newLocality)

		//process 
		res, err := h.sv.Create(locality)
		if err != nil {

			code, msg := eh.HandleError(err)
			response.Error(w, code, msg)
			return
		}
		
		//mapping 
		var data = helper.LocalityToLocalityDataResponseDTO(res)

		//Response
		response.JSON(w, http.StatusCreated, map[string]any{
			"data" : data,
		})
	}
}
