package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ISeller interface {
	GetAll() (map[int]model.Seller, error)
	GetByID(id int) (model.Seller, error)
}
