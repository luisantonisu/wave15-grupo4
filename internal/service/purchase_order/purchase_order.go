package service

import (
	"github.com/luisantonisu/wave15-grupo4/internal/domain/model"
	repository "github.com/luisantonisu/wave15-grupo4/internal/repository/purchase_order"
)

type PurchaseOrderService struct {
	repository repository.IPurchaseOrder
}

func NewPurchaseOrderService(repository repository.IPurchaseOrder) *PurchaseOrderService {
	return &PurchaseOrderService{
		repository: repository,
	}
}

// Create new purchase order
func (s *PurchaseOrderService) Create(purchaseOrder model.PurchaseOrderAttributes) (model.PurchaseOrder, error) {
	return s.repository.Create(purchaseOrder)
}
