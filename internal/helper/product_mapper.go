package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func ProductToProductResponseDTO(product map[int]model.Product) map[int]dto.ProductResponseDTO {
	data := make(map[int]dto.ProductResponseDTO)
	for _, value := range product {
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

func ProductRequestDTOToProduct(product dto.ProductRequestDTO) model.ProductAtrributes {
	data := model.ProductAtrributes{
		ProductCode:                    product.ProductCode,
		Description:                    product.Description,
		Width:                          product.Width,
		Height:                         product.Height,
		Length:                         product.Length,
		NetWeight:                      product.NetWeight,
		ExpirationRate:                 product.ExpirationRate,
		RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		FreezingRate:                   product.FreezingRate,
		ProductTypeID:                  product.ProductTypeId,
		SellerID:                       product.SellerId,
	}
	return data
}

func ProductRequestDTOPtrToProductPtr(product dto.ProductRequestDTOPtr) model.ProductAtrributesPtr {
	data := model.ProductAtrributesPtr{
		ProductCode:                    product.ProductCode,
		Description:                    product.Description,
		Width:                          product.Width,
		Height:                         product.Height,
		Length:                         product.Length,
		NetWeight:                      product.NetWeight,
		ExpirationRate:                 product.ExpirationRate,
		RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
		FreezingRate:                   product.FreezingRate,
		ProductTypeID:                  product.ProductTypeId,
		SellerID:                       product.SellerId,
	}
	return data
}
