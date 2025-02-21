package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func ProductRecordToProductRecordResponseDTO(product map[int]model.ProductRecord) map[int]dto.ProductRecordResponseDTO {
	data := make(map[int]dto.ProductRecordResponseDTO)
	for _, value := range product {
		data[value.ID] = dto.ProductRecordResponseDTO{
			ID: value.ID,
			ProductRecordRequestDTO: dto.ProductRecordRequestDTO{
				LastUpdateDate: value.LastUpdateDate,
				PurchasePrice:  value.PurchasePrice,
				SalePrice:      value.SalePrice,
				ProductId:      value.ProductId,
			},
		}
	}
	return data
}

func ProductRecordRequestDTOToProductRecord(product dto.ProductRecordRequestDTO) model.ProductRecordAtrributes {
	data := model.ProductRecordAtrributes{
		LastUpdateDate: product.LastUpdateDate,
		PurchasePrice:  product.PurchasePrice,
		SalePrice:      product.SalePrice,
		ProductId:      product.ProductId,
	}
	return data
}

func ProductRecordRequestDTOPtrToProductPtr(productRecord dto.ProductRecordRequestDTO) model.ProductRecordAtrributes {
	data := model.ProductRecordAtrributes{
		LastUpdateDate: productRecord.LastUpdateDate,
		PurchasePrice:  productRecord.PurchasePrice,
		SalePrice:      productRecord.SalePrice,
		ProductId:      productRecord.ProductId,
	}
	return data
}
