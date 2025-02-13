package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func BuyerRequestDTOToBuyer(buyerRequestDto dto.BuyerRequestDTO) model.Buyer {
	data := model.Buyer{
		BuyerAttributes: model.BuyerAttributes{
			CardNumberId: buyerRequestDto.CardNumberId,
			FirstName:    buyerRequestDto.FirstName,
			LastName:     buyerRequestDto.LastName,
		},
	}
	return data
}

func BuyerToBuyerResponseDTO(buyer model.Buyer) dto.BuyerResponseDTO {
	data := dto.BuyerResponseDTO{
		ID:           buyer.ID,
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}
	return data
}

func BuyerRequestDTOToBuyerAttributes(buyerRuquestDto dto.BuyerRequestDTO) model.BuyerAttributes {
	data := model.BuyerAttributes{
		CardNumberId: buyerRuquestDto.CardNumberId,
		FirstName:    buyerRuquestDto.FirstName,
		LastName:     buyerRuquestDto.LastName,
	}
	return data
}

func BuyerRequestDTOPtrToBuyerPtr(buyerRequestDto dto.BuyerRequestDTOPtr) model.BuyerAttributesPtr {
	data := model.BuyerAttributesPtr{
		CardNumberId: buyerRequestDto.CardNumberId,
		FirstName:    buyerRequestDto.FirstName,
		LastName:     buyerRequestDto.LastName,
	}
	return data
}
