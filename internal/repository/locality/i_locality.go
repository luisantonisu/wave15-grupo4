package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ILocality interface {
	GetByID(id int) (model.LocalityDBModel, error)
	Create(locality model.LocalityDBModel) (model.LocalityDBModel, error)
	CarriersReport(id *int) ([]model.CarriersReport, error)
	SellersReport(id *int) ([]model.LocalityReport, error)
}
