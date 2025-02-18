package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ISeller interface {
	GetAll() ([]model.Seller, error)
	GetByID(id int) (model.Seller, error)
	Create(seller model.SellerAttributes) (model.Seller, error)
	Update(id int, seller model.SellerAttributes) (model.Seller, error)
	Delete(id int) error
}
