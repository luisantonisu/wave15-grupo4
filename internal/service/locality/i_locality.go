package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ILocality interface {
	Create(locality model.Locality) (model.Locality, error)
	CarriersReport(id *int) ([]model.CarriersReport, error)
	SellersReport(id *int) ([]model.LocalityReport, error)
}
