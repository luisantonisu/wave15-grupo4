package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IProductBatch interface {
	Create(productBatch model.ProductBatch) (model.ProductBatch, error)
}
