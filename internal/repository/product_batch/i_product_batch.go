package repsotory

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IProductBatch interface {
	Create(productBatch model.ProductBatchAttributes) (model.ProductBatch, error)
}
