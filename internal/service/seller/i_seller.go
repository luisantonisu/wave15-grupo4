package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ISeller interface {
	GetAll() (sellers map[int]model.Seller, err error)
	GetByID(id int) (seller model.Seller, err error)
}