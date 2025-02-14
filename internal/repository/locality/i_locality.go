package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"


type ILocality interface {
	Create(locality model.LocalityDBModel) (model.LocalityDBModel, error)
}