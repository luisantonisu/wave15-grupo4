package service

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type ICarry interface {
	Create(carry model.Carry) (model.Carry, error)
}
