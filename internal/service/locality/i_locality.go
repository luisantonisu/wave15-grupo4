package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ILocality interface {
	Create(locality model.Locality) (model.Locality, error)
	Report(id int) (map[int]model.CarriersByLocalityReport, error)
	SellersReport(id *int) ([]model.LocalityReport, error)
}
