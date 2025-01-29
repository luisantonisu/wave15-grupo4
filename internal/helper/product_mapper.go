package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func ProductToProductResponseDTO(employees map[int]model.Product) map[int]dto.ProductResponseDTO {
	data := make(map[int]dto.ProductResponseDTO)
	for _, value := range employees {
		data[value.ID] = dto.ProductResponseDTO{
			ID:                             value.ID,
			ProductCode:                    value.ProductAtrributes.ProductCode,
			Description:                    value.ProductAtrributes.Description,
			Width:                          value.ProductAtrributes.Width,
			Height:                         value.ProductAtrributes.Height,
			Length:                         value.ProductAtrributes.Length,
			NetWeight:                      value.ProductAtrributes.NetWeight,
			ExpirationRate:                 value.ProductAtrributes.ExpirationRate,
			RecommendedFreezingTemperature: value.ProductAtrributes.RecommendedFreezingTemperature,
			FreezingRate:                   value.ProductAtrributes.FreezingRate,
			ProductTypeId:                  value.ProductAtrributes.ProductTypeID,
			SellerId:                       value.ProductAtrributes.SellerID,
		}
	}
	return data
}

func ProductRequestDTOToProduct(employees dto.ProductRequestDTO) model.ProductAtrributes {
	data := model.ProductAtrributes{
		ProductCode:                    employees.ProductCode,
		Description:                    employees.Description,
		Width:                          employees.Width,
		Height:                         employees.Height,
		Length:                         employees.Length,
		NetWeight:                      employees.NetWeight,
		ExpirationRate:                 employees.ExpirationRate,
		RecommendedFreezingTemperature: employees.RecommendedFreezingTemperature,
		FreezingRate:                   employees.FreezingRate,
		ProductTypeID:                  employees.ProductTypeId,
		SellerID:                       employees.SellerId,
	}
	return data
}
