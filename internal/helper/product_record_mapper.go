package helper

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/dto"
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
)

func ProductRecordToProductRecordResponseDTO(product model.ProductRecord) dto.ProductRecordResponseDTO {
	data := dto.ProductRecordResponseDTO{
		ID: product.ID,
		ProductRecordRequestDTO: dto.ProductRecordRequestDTO{
			LastUpdateDate: product.LastUpdateDate,
			PurchasePrice:  product.PurchasePrice,
			SalePrice:      product.SalePrice,
			ProductId:      product.ProductId,
		},
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
