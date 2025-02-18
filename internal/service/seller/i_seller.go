package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ISeller interface {
	GetAll() (sellers []model.Seller, err error)
	GetByID(id int) (seller model.Seller, err error)
	Create(seller model.Seller) (model.Seller, error)
	Update(id int, seller model.SellerAttributes) (model.Seller, error)
	Delete(id int) error
}
