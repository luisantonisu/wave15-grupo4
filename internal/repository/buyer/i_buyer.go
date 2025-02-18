package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IBuyer interface {
	Create(buyer model.BuyerAttributes) (model.Buyer, error)
	GetAll() ([]model.Buyer, error)
	GetByID(id int) (model.Buyer, error)
	Delete(id int) error
	Update(id int, attributes model.BuyerAttributes) (model.Buyer, error)
	PurchaseOrderReport(id *int) ([]model.ReportPurchaseOrders, error)
}
