package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func ProductToProductResponseDTO(product model.Product) dto.ProductResponseDTO {
	data := dto.ProductResponseDTO{
		ID: product.ID,
		ProductRequestDTO: dto.ProductRequestDTO{
			ProductCode:                    product.ProductCode,
			Description:                    product.Description,
			Width:                          product.Width,
			Height:                         product.Height,
			Length:                         product.Length,
			NetWeight:                      product.NetWeight,
			ExpirationRate:                 product.ExpirationRate,
			RecommendedFreezingTemperature: product.RecommendedFreezingTemperature,
			FreezingRate:                   product.FreezingRate,
			ProductTypeId:                  product.ProductTypeID,
			SellerId:                       product.SellerID,
		},
	}
	return data
}

func ProductRequestDTOToProduct(product dto.ProductRequestDTO) model.ProductAttributes {
	data := model.ProductAttributes{
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

func ProductRecordCountToProductRecordCountResponseDTO(product model.ProductRecordCount) dto.ProductRecordCountResponseDTO {
	data := dto.ProductRecordCountResponseDTO{
		ProductID:   product.ProductID,
		Description: product.Description,
		Count:       product.Count,
	}
	return data
}
