package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func BuyerDtoToBuyer(buyerDto dto.BuyerDTO) model.Buyer {
	data := model.Buyer{
		Id: buyerDto.Id,
		BuyerAttributes: model.BuyerAttributes{
			CardNumberId: buyerDto.CardNumberId,
			FirstName:    buyerDto.FirstName,
			LastName:     buyerDto.LastName,
		},
	}
	return data
}

func BuyerToBuyerDto(buyer model.Buyer) dto.BuyerDTO {
	data := dto.BuyerDTO{
		Id:           buyer.Id,
		CardNumberId: buyer.CardNumberId,
		FirstName:    buyer.FirstName,
		LastName:     buyer.LastName,
	}
	return data
}
