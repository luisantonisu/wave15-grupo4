package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ILocality interface {
	GetByID(id int) (model.LocalityDBModel, error)
	Create(locality model.LocalityDBModel) (model.LocalityDBModel, error)
	Report(id int) (map[int]model.CarriersByLocalityReport, error) 
	SellersReport(id *int) (map[int]model.LocalityReport, error)
}
