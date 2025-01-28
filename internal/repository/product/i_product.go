package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IProduct interface {
	GetProduct() (prMap map[int]model.Product, err error)
}
