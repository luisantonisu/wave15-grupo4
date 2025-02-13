package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func SellerToSellerResponseDTO(seller model.Seller) dto.SellerResponseDTO {
	data := dto.SellerResponseDTO{
		ID:          seller.ID,
		CompanyID:   seller.CompanyID,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
	}

	return data
}

func SellerRequestDTOToSeller(sellerRequestDTO dto.SellerRequestDTO) model.Seller {
	data := model.Seller{
		SellerAttributes: model.SellerAttributes{
			CompanyID:   sellerRequestDTO.CompanyID,
			CompanyName: sellerRequestDTO.CompanyName,
			Address:     sellerRequestDTO.Address,
			Telephone:   sellerRequestDTO.Telephone,
		},
	}

	return data
}

func SellerRequestDTOPtrToSellerPtr(seller dto.SellerRequestDTOPtr) model.SellerAttributesPtr {
	data := model.SellerAttributesPtr{
		CompanyID:   seller.CompanyID,
		CompanyName: seller.CompanyName,
		Address:     seller.Address,
		Telephone:   seller.Telephone,
	}
	return data
}
