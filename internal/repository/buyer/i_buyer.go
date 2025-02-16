package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IBuyer interface {
	Create(buyer model.BuyerAttributes) (model.Buyer, error)
	GetAll() (map[int]model.Buyer, error)
	GetByID(id int) (model.Buyer, error)
	Delete(id int) error
	Update(id int, attributes model.BuyerAttributesPtr) (model.Buyer, error)
	Report(id int) (map[int]model.ReportPurchaseOrders, error)
}
