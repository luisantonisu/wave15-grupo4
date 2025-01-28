package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func MapProductToProductDTO(employees map[int]model.Product) map[int]dto.ProductDTO {
	data := make(map[int]dto.ProductDTO)
	for _, value := range employees {
		data[value.ID] = dto.ProductDTO{
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
			ProductTypeId:                  value.ProductAtrributes.ProductTypeId,
			SellerId:                       value.ProductAtrributes.SellerId,
		}
	}
	return data
}

func MapProductDTOToProduct(employees dto.ProductDTO) model.Product {
	data := model.Product{
		ID: employees.ID,
		ProductAtrributes: model.ProductAtrributes{
			ProductCode:                    employees.ProductCode,
			Description:                    employees.Description,
			Width:                          employees.Width,
			Height:                         employees.Height,
			Length:                         employees.Length,
			NetWeight:                      employees.NetWeight,
			ExpirationRate:                 employees.ExpirationRate,
			RecommendedFreezingTemperature: employees.RecommendedFreezingTemperature,
			FreezingRate:                   employees.FreezingRate,
			ProductTypeId:                  employees.ProductTypeId,
			SellerId:                       employees.SellerId,
		},
	}
	return data
}
