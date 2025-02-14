package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func CarryToCarryResponseDTO(carry model.Carry) dto.CarryResponseDTO {
	data := dto.CarryResponseDTO{
		ID:          carry.ID,
		CarryID:     carry.CarryID,
		CompanyName: carry.CompanyName,
		Address:     carry.Address,
		Telephone:   carry.Telephone,
		LocalityID:  carry.LocalityID,
	}
	return data
}

func CarryRequestDTOToCarry(carryRequestDTO dto.CarryRequestDTO) model.Carry {
	data := model.Carry{
		CarryAttributes: model.CarryAttributes{
			CarryID:     carryRequestDTO.CarryID,
			CompanyName: carryRequestDTO.CompanyName,
			Address:     carryRequestDTO.Address,
			Telephone:   carryRequestDTO.Telephone,
			LocalityID:  carryRequestDTO.LocalityID,
		},
	}
	return data
}

func CarryRequestDTOPtrToCarryAttributesPtr(carryRequestDTOPtr dto.CarryRequestDTOPtr) model.CarryAttributesPtr {
	data := model.CarryAttributesPtr{
		CarryID:     carryRequestDTOPtr.CarryID,
		CompanyName: carryRequestDTOPtr.CompanyName,
		Address:     carryRequestDTOPtr.Address,
		Telephone:   carryRequestDTOPtr.Telephone,
		LocalityID:  carryRequestDTOPtr.LocalityID,
	}
	return data
}
