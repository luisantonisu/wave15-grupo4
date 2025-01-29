package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func AdaptSellersListToSellerListDTO(sellers map[int]model.Seller) map[int]dto.SellerDTO {
	adapted := make(map[int]dto.SellerDTO)
	for _, value := range sellers {
		adapted[value.Id] = dto.SellerDTO{
			Id : value.Id,
			CompanyId: value.CompanyId,
			CompanyName: value.CompanyName,
			Address: value.Address,
			Telephone: value.Telephone,
		}
	}
	return adapted
}
