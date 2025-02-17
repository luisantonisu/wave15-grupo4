package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ILocality interface {
	Create(locality model.Locality) (model.Locality, error)
	SellersReport(id *int) (map[int]model.LocalityReport, error)
}
