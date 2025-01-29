package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IBuyer interface {
	Create(model.Buyer) (model.Buyer, error)
}
