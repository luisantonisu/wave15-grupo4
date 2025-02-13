package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func ProductToProductResponseDTO(product model.Product) dto.ProductResponseDTO {
	data := dto.ProductResponseDTO{
		ID:                             product.ID,
		ProductCode:                    product.ProductAtrributes.ProductCode,
		Description:                    product.ProductAtrributes.Description,
		Width:                          product.ProductAtrributes.Width,
		Height:                         product.ProductAtrributes.Height,
		Length:                         product.ProductAtrributes.Length,
		NetWeight:                      product.ProductAtrributes.NetWeight,
		ExpirationRate:                 product.ProductAtrributes.ExpirationRate,
		RecommendedFreezingTemperature: product.ProductAtrributes.RecommendedFreezingTemperature,
		FreezingRate:                   product.ProductAtrributes.FreezingRate,
		ProductTypeId:                  product.ProductAtrributes.ProductTypeID,
		SellerId:                       product.ProductAtrributes.SellerID,
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

func ProductRecordCountToProductRecordCountResponseDTO(product model.ProductRecordCount) dto.ProductRecordCountResponseDTO {
	data := dto.ProductRecordCountResponseDTO{
		ProductID:   product.ProductID,
		Description: product.Description,
		Count:       product.Count,
	}
	return data
}
