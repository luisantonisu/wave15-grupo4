package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ISeller interface {
	GetAll() (sellers map[int]model.Seller, err error)
}
