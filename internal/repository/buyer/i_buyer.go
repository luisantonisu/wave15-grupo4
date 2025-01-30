package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IBuyer interface {
	Create(model.Buyer) (model.Buyer, error)
	GetAll() (map[int]model.Buyer, error)
	GetByID(int) (model.Buyer, error)
	Delete(int) error
	Update(int, model.BuyerAttributesPtr) (model.Buyer, error)
}
