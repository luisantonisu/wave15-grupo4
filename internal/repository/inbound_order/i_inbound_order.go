package repository

import "github.com/luisantonisu/wave15-grupo4/internal/domain/model"

type IInboundOrder interface {
	AlreadyExists(attribute string, value int) bool
	CreateInboundOrder(inboundOrder model.InboundOrderAttributes) (model.InboundOrder, error)
}
