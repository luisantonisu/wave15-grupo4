package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IPurchaseOrder interface {
	Create(purchaseOrder model.PurchaseOrderAttributes) (model.PurchaseOrder, error)
}
