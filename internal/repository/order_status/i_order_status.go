package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IOrderStatus interface {
	GetByID(id int) (model.OrderStatus, error)
}
